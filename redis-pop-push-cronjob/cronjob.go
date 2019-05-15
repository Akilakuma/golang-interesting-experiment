package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	cron "gopkg.in/robfig/cron.v2"
)

// CronJob 背景排程
type CronJob struct {
	// 背景名稱
	Name string
	// 執行週期
	Spec string
	// 執行工作
	Cmd func() error
	// 是否可以重複
	IsOverLapping bool
	// EntryID
	entryID cron.EntryID
	// 正在執行
	running bool
	// 讀寫鎖
	mux *sync.RWMutex
	// 等待通道
	wg *sync.WaitGroup
}

// RunJob 設定工作
func RunJob(jobs []*CronJob) {
	bg := cron.New()

	for _, job := range jobs {
		job.Init()
		pid, err := bg.AddJob(job.Spec, job)
		if err != nil {
			log.Fatalln(err)
		}
		job.SetEntryID(pid)
	}

	// 開始排程
	log.Println("INFO", `
	🐳  啟動排程囉~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ 🐳
	`)
	bg.Start()

	// 等待結束訊號
	<-gracefulShutdown()
	log.Println("WARNIGN", `
	🚦  收到訊號囉~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ 🚦
	`)

	// 停止排程
	bg.Stop()

	// 等待背景結束
	for _, job := range jobs {
		job.Wait()
	}

	log.Println("INFO", `
	🔥  結束囉~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ 🔥
	`)

}

func gracefulShutdown() (sig chan os.Signal) {
	sig = make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	return
}

// Run 執行
func (c *CronJob) Run() {

	if c.IsOverLapping {
		c.wg.Add(1)
		c.Cmd()
		c.wg.Done()
		return
	}

	if c.running {
		return
	}

	// 執行背景
	c.wg.Add(1)
	c.mux.Lock()
	c.running = true
	c.mux.Unlock()

	c.Cmd()

	c.mux.Lock()
	c.running = false
	c.mux.Unlock()
	c.wg.Done()

}

// SetEntryID 設定entryID
func (c *CronJob) SetEntryID(id cron.EntryID) {
	c.entryID = id
}

// Init 初始化
func (c *CronJob) Init() {
	c.mux = new(sync.RWMutex)
	c.wg = new(sync.WaitGroup)
}

// Wait 等待結束
func (c *CronJob) Wait() {
	c.wg.Wait()
}

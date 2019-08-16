package main

import (
	"log"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

var (
	successCount int
	errCount     int
	m            = new(sync.RWMutex)
)

// TimeMsg timer定時回報
func TimeMsg(jb chan struct{}, wg *sync.WaitGroup) {

	d := time.Duration(time.Second * 10)
	t := time.NewTicker(d)
	defer t.Stop()

	var trigger struct{}

	for {
		<-t.C
		// log.Println(runtime.NumGoroutine())
		for i := 0; i < 500; i++ {
			jb <- trigger
		}
		wg.Wait()
		log.Println("done")
	}
}

func main() {

	var (
		wg sync.WaitGroup
	)

	jobChan := make(chan struct{}, 200)
	// fireChan := make(chan struct{}, 200)
	for i := 0; i < 500; i++ {
		go worker(jobChan, &wg)
	}

	go TimeMsg(jobChan, &wg)
	for {
	}
}

func worker(jobs <-chan struct{}, wg *sync.WaitGroup) {
	// log.Println("worker gogo")
	for j := range jobs {
		// log.Println("hi job", j)
		wg.Add(1)
		fire(wg, j)
	}
}

func fire(wg *sync.WaitGroup, j struct{}) {

	_, dialErr := jsonrpc.Dial("tcp", "127.0.0.1:7777")
	if dialErr != nil {
		if dialErr.Error() == "dial tcp 127.0.0.1:7777: connect: connection reset by peer" {
			// log.Println("hello")

			m.Lock()
			errCount++
			m.Unlock()
		}
		log.Println("msg", dialErr)
	} else {
		m.Lock()
		successCount++
		m.Unlock()
	}

	wg.Done()
}

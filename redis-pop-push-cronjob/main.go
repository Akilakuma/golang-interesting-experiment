package main

import (
	"log"
	"os"
	"strconv"
	"time"
	// "reflect"
	"github.com/gomodule/redigo/redis"
)

func main() {

	// 打開redis

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "push":
			// 有時間讓pop起起來
			time.Sleep(10)
			SetPushByCron()
		case "pop":
			SetPopByCron()
		}
	}

}

// PushRedis 對redis 做rpsuh
func PushRedis() error {



	conn := getConnectionPool()
	redisC := conn.Get()
	defer redisC.Close()

	// 結果寫回array
	for i := 0; i < 5000; i++ {

		_, err := redisC.Do("RPUSH", "data_queue", "num_"+strconv.Itoa(i))
		if err != nil {
			log.Println(err.Error())
		}
	}

	return nil
}

// PopRedis 對redis 做lpop
func PopRedis() error {

	conn := getConnectionPool()
	redisC := conn.Get()
	defer redisC.Close()

	var (
		before   string
		now      string
		dataList []string
	)

	// 先數一個長度
	redisLen, lenErr := redis.Int(redisC.Do("llen", "data_queue"))
	if lenErr != nil {
		if lenErr.Error() != "redigo: nil returned" {
			log.Println(lenErr.Error())
		}
	}

	// 一筆一筆pop 出來
	for redisLen > 0 {
		reversalList, popErr := redis.String(redisC.Do("LPOP", "data_queue"))
		if popErr != nil {
			if popErr.Error() != "redigo: nil returned" {
				log.Println(popErr.Error())
			}
		}

		dataList = append(dataList, reversalList)
		redisLen--
	}

	// 結果寫回array
	for _, v := range dataList {

		log.Println(v)
		now = v
		if before == now {
			log.Println("有出現pop出同樣的值")
			panic("有出現pop出同樣的值")
		}

		before = now
	}


	return nil
}

// getConnectionPool 實質建立連線
func getConnectionPool() *redis.Pool {

	return &redis.Pool{
		Wait:        true,
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Println(err.Error())
			}
			return err
		},
	}

}

// SetPushByCron 啟動push 的 cronjob
func SetPushByCron() {
	// 排程清單
	jobs := []*CronJob{
		{Name: "push", Spec: "*/2 * * * * *", Cmd: PushRedis, IsOverLapping: false}, // 定時 push 到 redis
	}

	// 啟動排程
	RunJob(jobs)
}

// SetPopByCron 啟動pop 的 cronjob
func SetPopByCron() {
	// 排程清單
	jobs := []*CronJob{
		{Name: "pop", Spec: "*/1 * * * * *", Cmd: PopRedis, IsOverLapping: false}, // 定時 從 redis pop
	}

	// 啟動排程
	RunJob(jobs)
}

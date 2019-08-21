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

func main() {

	var (
		wg sync.WaitGroup
	)

	// 總共一萬次
	for i := 0; i < 1; i++ {

		timeS := time.Now()

		// 每次起200個gorutine去下語法
		for j := 0; j < 1000; j++ {
			wg.Add(1)
			go test1(&wg)

		}

		wg.Wait()

		timeE := time.Now().Sub(timeS)
		log.Println("timeE", timeE)

		time.Sleep(1000 * time.Millisecond)
		log.Println("errCount", errCount)
		log.Println("successCount", successCount)
		if i%10 == 0 {
			log.Println(i, "輪結束＝＝＝＝＝＝＝＝＝＝")
		}
	}
	log.Println("結束")
}

func test1(wg *sync.WaitGroup) {

	// nowBeforeDial := time.Now().UTC().Add(8 * time.Hour)

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

	// nowAfterDial := time.Now().UTC().Add(8 * time.Hour)
	// b := nowAfterDial.Sub(nowBeforeDial)

	// if b > 100*time.Millisecond {
	// 	// log.Println(b)
	// }

	wg.Done()
}

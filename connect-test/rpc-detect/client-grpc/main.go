package main

import (
	"context"
	"log"
	"rpcdetect/pb"
	"runtime"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {

	log.Println(runtime.Version())

	var (
		wg sync.WaitGroup
	)

	// 總共一萬次
	for i := 0; i < 1000; i++ {

		// 每次起200個gorutine去下語法
		for j := 0; j <1000; j++ {
			wg.Add(1)
			go test1(&wg)

		}

		wg.Wait()
		time.Sleep(1000 * time.Millisecond)
		if i%10 == 0 {
			log.Println(i, "輪結束＝＝＝＝＝＝＝＝＝＝")
		}
	}
	log.Println("結束")
}

func test1(wg *sync.WaitGroup) {

	nowBeforeDial := time.Now().UTC().Add(8 * time.Hour)
	now := nowBeforeDial.Format("2006-01-02 15:04:05")

	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()
	nowAfterDial := time.Now().UTC().Add(8 * time.Hour)
	b := nowAfterDial.Sub(nowBeforeDial)

	if b > 100*time.Millisecond {
		log.Println(b)
	}

	aService := pb.NewConnectClient(conn)
	pingPara := &pb.ConnRequest{Rand: 1, TimeAt: now}

	_, err = aService.Ping(context.Background(), pingPara)
	if err != nil {
		log.Println("有錯誤發生:" + err.Error())
	}
	// log.Println("res.Result:", res.Result)

	wg.Done()
}

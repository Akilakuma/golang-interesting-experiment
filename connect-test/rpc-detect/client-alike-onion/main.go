package main

import (
	"context"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"runtime"
	"sync"
	"time"
)

func main() {

	var (
		wg sync.WaitGroup
	)

	for {
		for i := 0; i < 1000; i++ {

			go curlrpc(&wg)
		}
		time.Sleep(1 * time.Second)
		log.Println("gorutine num is ", runtime.NumGoroutine())
	}

}

func curlrpc(wg *sync.WaitGroup) {
	var (
		timeoutDuration = time.Second * 1
		ctx             context.Context
		cancel          context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.Background(), timeoutDuration)

	go func() {
		defer cancel()
		client, dialErr := jsonrpc.Dial("tcp", "127.0.0.1:7777")
		if dialErr != nil {
			log.Println("rpc 連線失敗", dialErr)
			ctx = context.WithValue(ctx, "Error Key", dialErr)

			return
		}

		select {
		case <-ctx.Done():
			log.Println("RPC 連線成功但是timeout")
			client.Close()
		default:
			ctx = context.WithValue(ctx, "Client Key", client)
		}
	}()
	runtime.Gosched()

	<-ctx.Done()

	ctxErr := ctx.Err()
	var client *rpc.Client
	switch ctxErr {
	case context.Canceled:
		dialErr := ctx.Value("Error Key")
		if dialErr != nil {
			log.Println(dialErr)
			return
		}

		conn := ctx.Value("Client Key")
		if conn != nil {
			client = conn.(*rpc.Client)
			defer client.Close()
		}
	case context.DeadlineExceeded:
		log.Println("已經是deadline")
		return
	}

}

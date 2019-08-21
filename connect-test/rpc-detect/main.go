package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	log.Println()
	go func(c context.Context) {

		defer cancel()
		time.Sleep(2 * time.Second)

		select {
		case <-c.Done():
			log.Println("context is done")
		default:
			log.Println("income default")
		}
	}(ctx)

	// runtime.Gosched()

	// time.Sleep(3 * time.Second)

	<-ctx.Done()
	log.Println("keep continue!")

	for {
	}
}

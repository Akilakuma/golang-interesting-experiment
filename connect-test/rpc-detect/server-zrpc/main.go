package main

import (
	"context"
	"log"
	"connect-test/rpc-detect/pb"
	"time"

	"github.com/yam8511/zrpc"
)

type ConnServer struct{}

func (conn *ConnServer) Ping(c context.Context, p *pb.ConnRequest) (*pb.ConnResponse, error) {
	// log.Println("rand：", p.Rand)
	// log.Println("time", p.TimeAt)

	_, conErr := time.Parse("2006-01-02 15:04:05", p.TimeAt)
	if conErr != nil {
		log.Println(conErr)
	}
	// timeNow := time.Now().UTC().Add(8 * time.Hour)
	// b:= timeNow.Sub(timeIncome)
	// log.Println(b)

	return &pb.ConnResponse{
		Result: 1,
	}, nil
}

func main() {

	server := zrpc.NewServer()
	server.RegisterName("ping", new(ConnServer))

	server.SetJSONRPCAddress("127.0.0.1:7777")
	err := server.Listen()
	if err != nil {
		log.Println(err)
	}
}

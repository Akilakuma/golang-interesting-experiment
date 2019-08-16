package main

import (
	"context"
	"log"
	"net/http"
	"net/rpc"
	"rpcdetect/pb"
	"time"
)

type ConnServer struct{}

func (conn *ConnServer) Ping(c context.Context, p *pb.ConnRequest) (*pb.ConnResponse, error) {

	_, conErr := time.Parse("2006-01-02 15:04:05", p.TimeAt)
	if conErr != nil {
		log.Println(conErr)
	}

	return &pb.ConnResponse{
		Result: 1,
	}, nil
}

func main() {

	rpc.NewServer()
	rpc.Register(new(ConnServer))
	rpc.HandleHTTP()
	if err := http.ListenAndServe("127.0.0.1:7777", nil); err != nil {
		log.Println(err)
	}

}

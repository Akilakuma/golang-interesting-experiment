package main

import (
	"context"
	"log"
	"net"
	"rpcdetect/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// "github.com/yam8511/zrpc"
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
	// server := zrpc.NewServer()
	// server.SetHTTPAddress("127.0.0.1:6666")
	// server.SetJSONRPCAddress("127.0.0.1:6667")
	// err := server.Listen()
	// if err != nil {
	// 	log.Println(err)
	// }

	log.Println("grpc 開始啟動")
	go grpc.TimeMsg()

	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Println(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterConnectServer(s, &ConnServer{})

	reflection.Register(s)

	log.Println("grpc 啟動成功")

	if err := s.Serve(listen); err != nil {
		log.Println(err.Error())
	}

}

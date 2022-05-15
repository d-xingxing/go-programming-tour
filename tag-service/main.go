package main

import (
	"flag"
	pb "github.com/go-programming-tour/tag-service/proto"
	"github.com/go-programming-tour/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	// 双端口各自监听http和rpc协议
	//grpcPort string
	//httpPort string

	// 在同端口监听HTTP
	port string
)

// 自定义错误
type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message, omitempty"`
}

func init() {
	//flag.StringVar(&httpPort, "http_port", "9001", "HTTP启动端口号")
	//flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC启动端口号")

	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.Parse()
}

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}

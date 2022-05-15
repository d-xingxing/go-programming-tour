package main

import (
	"flag"
	pb "github.com/go-programming-tour/tag-service/proto"
	"github.com/go-programming-tour/tag-service/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var (
	// 双端口各自监听http和rpc协议
	grpcPort string
	httpPort string

	// 在同端口监听HTTP
	port string
)

// 自定义错误
type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message, omitempty"`
}

func init() {
	flag.StringVar(&httpPort, "http_port", "9001", "HTTP启动端口号")
	flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC启动端口号")

	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.Parse()
}

func main() {
	//serverVersion1()
	serverVersion2()

}

// 版本2------------------同端口监听HTTP------------------------------------
//在 Go 语言中，我们可以使用第三方开源库 cmux 来实现多协议支持的功能，cmux 是根据有效负载（payload）
//对连接进行多路复用（也就是匹配连接的头几个字节来进行区分当前连接的类型），可以在同一 TCP Listener 上
//提供 gRPC、SSH、HTTPS、HTTP、Go RPC 以及几乎所有其它协议的服务，是一个相对通用的方案。
func serverVersion2() {
	l, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP Server err: %v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer()
	httpS := RunHttpServer(port)
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}

func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

// 版本1-----------------------另起端口监听------------------------------
//func serverVersion1() {
//	// 版本一：另起端口监听 HTTP
//	errs := make(chan error)
//
//	go func() {
//		err := RunHttpServer(httpPort)
//		if err != nil {
//			errs <- err
//		}
//	}()
//	go func() {
//		err := RunGrpcServer(grpcPort)
//		if err != nil {
//			errs <- err
//		}
//	}()
//
//	select {
//	case err := <-errs:
//		log.Fatalf("Run server err: %v", err)
//	}
//}

// 将原本的gRPC服务启动端口调整为对HTTP1.1 和 gRPC端口号的读取
// 针对HTTP的RunHttpServer方法,可用于基本的心跳检测
//func RunHttpServer(port string) error {
//	serverMux := http.NewServeMux() // 初始化一个多路复用器
//	serverMux.HandleFunc("/ping",
//		func(w http.ResponseWriter, r *http.Request) {
//			_, _ = w.Write([]byte(`pong`))
//		},
//	)
//	return http.ListenAndServe(":"+port, serverMux)
//}
//
//func RunGrpcServer(port string) error {
//	s := grpc.NewServer()
//	pb.RegisterTagServiceServer(s, server.NewTagServer())
//	reflection.Register(s) // 使用工具grocurl的前提是gRPC server已经注册了反射服务
//	lis, err := net.Listen("tcp", ":"+port)
//	if err != nil {
//		return err
//	}
//	return s.Serve(lis)
//}

package main

import (
	"context"
	pb "github.com/go-programming-tour/tag-service/proto"
	"google.golang.org/grpc"
	"log"
)

// 编写示例来调用gRPC服务
func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8003", nil)
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Rust"})

	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...) // DialContext()并不能立刻创建连接,因为其是非阻塞的
}

// Auth 实现Auth认证
type Auth struct {
	AppKey    string
	AppSecret string
}

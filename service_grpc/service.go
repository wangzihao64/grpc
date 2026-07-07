package main

import (
	"context"
	"fmt"
	"grpcdemo/grpc_proto/hello_grpc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type HelloService struct {
	hello_grpc.UnimplementedHelloServiceServer
}

func (HelloService) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (res *hello_grpc.HelloReponse, err error) {
	fmt.Println("HelloService.SayHello")
	fmt.Println(request.Message)
	return &hello_grpc.HelloReponse{
		Name:    "王子豪",
		Message: "ok",
	}, nil
}
func main() {
	//监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	//创建一个grpc服务器实例
	s := grpc.NewServer()
	server := HelloService{}
	//将server结构体注册为gRPC服务
	hello_grpc.RegisterHelloServiceServer(s, &server)
	fmt.Println("grpc server runing :8080")
	//开始处理客户端请求
	err = s.Serve(listen)
}

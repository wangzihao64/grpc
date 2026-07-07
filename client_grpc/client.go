package main

import (
	"context"
	"fmt"
	"grpcdemo/grpc_proto/hello_grpc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := ":8080"
	//使用grpc.Dial创建一个到指定地址的gRPC连接
	//此处使用不安全的证书来实现SSL/TSL连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr[%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	//初始化客户端
	client := hello_grpc.NewHelloServiceClient(conn)
	response, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "grpcdemo",
		Message: "nihao",
	})
	fmt.Println(response, err)
}

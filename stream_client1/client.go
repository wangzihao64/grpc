package main

import (
	"context"
	"fmt"
	"grpcdemo/stream_proto/proto"
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
	client := proto.NewClientstreamClient(conn)
	stream, err := client.UploadFile(context.Background())
	for i := 0; i < 10; i++ {
		stream.Send(&proto.FileRequest{
			FileName: fmt.Sprintf("第%d次\n", i),
		})
	}
	response, err := stream.CloseAndRecv()
	fmt.Println(response, err)
}

package main

import (
	"context"
	"fmt"
	"grpcdemo/grpc_proto/duo_grpc"
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
	orderclient := duo_grpc.NewOrderServiceClient(conn)
	response, err := orderclient.Buy(context.Background(), &duo_grpc.Request{
		Name: "BuyClient",
	})
	fmt.Println(response)
	videoclient := duo_grpc.NewVideoServiceClient(conn)
	response, err = videoclient.Look(context.Background(), &duo_grpc.Request{
		Name: "BuyClient",
	})
	fmt.Println(response)
}

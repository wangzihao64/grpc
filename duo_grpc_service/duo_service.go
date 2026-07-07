package main

import (
	"context"
	"fmt"
	"grpcdemo/grpc_proto/duo_grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type VideoService struct {
	duo_grpc.UnimplementedVideoServiceServer
}

func (VideoService) Look(ctx context.Context, request *duo_grpc.Request) (response *duo_grpc.Response, err error) {
	fmt.Println(request)
	return &duo_grpc.Response{
		Name: "serverVideo",
	}, nil
}

type OrderService struct {
	duo_grpc.UnimplementedOrderServiceServer
}

func (OrderService) Buy(ctx context.Context, request *duo_grpc.Request) (response *duo_grpc.Response, err error) {
	fmt.Println(request)
	return &duo_grpc.Response{
		Name: "serverOrder",
	}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	grpcServer := grpc.NewServer()
	duo_grpc.RegisterVideoServiceServer(grpcServer, &VideoService{})
	duo_grpc.RegisterOrderServiceServer(grpcServer, &OrderService{})
	fmt.Println("grpc server listen on :8080")
	err = grpcServer.Serve(listen)
}

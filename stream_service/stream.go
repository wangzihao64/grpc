package main

import (
	"fmt"
	"grpcdemo/stream_proto/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ServiceStream struct {
	proto.UnimplementedServicestreamServer
}

func (ServiceStream) Fun(request *proto.Request, stream proto.Servicestream_FunServer) error {
	fmt.Println(request)
	for i := 0; i < 10; i++ {
		stream.Send(&proto.Response{
			Text: fmt.Sprintf("hello %d", i),
		})
	}
	return nil
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterServicestreamServer(server, &ServiceStream{})
	server.Serve(listen)
}

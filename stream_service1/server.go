package main

import (
	"fmt"
	"grpcdemo/stream_proto/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ClientStream struct {
	proto.UnimplementedClientstreamServer
}

func (ClientStream) UploadFile(stream proto.Clientstream_UploadFileServer) error {
	for i := 0; i < 10; i++ {
		response, err := stream.Recv()
		fmt.Println(response, err)
	}
	stream.SendAndClose(&proto.Response{Text: "完毕了"})
	return nil
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterClientstreamServer(server, &ClientStream{})
	server.Serve(listen)
}

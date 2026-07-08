package main

import (
	"fmt"
	"grpcdemo/stream_proto/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type ServiceStream struct {
	proto.UnimplementedServiceStreamServer
}

func (ServiceStream) Fun(request *proto.Request, stream proto.ServiceStream_FunServer) error {
	fmt.Println(request)
	for i := 0; i < 10; i++ {
		stream.Send(&proto.Response{
			Text: fmt.Sprintf("hello %d", i),
		})
	}
	return nil
}
func (ServiceStream) DownLoadFile(request *proto.Request, stream proto.ServiceStream_DownLoadFileServer) error {
	fmt.Println(request)
	file, err := os.Open("static/王子豪简历v13.pdf")
	if err != nil {
		return err
	}
	defer file.Close()
	for {
		buf := make([]byte, 1024)
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		fmt.Println("%s", buf)
		stream.Send(&proto.FileResponse{
			Content: buf,
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
	proto.RegisterServiceStreamServer(server, &ServiceStream{})
	server.Serve(listen)
}

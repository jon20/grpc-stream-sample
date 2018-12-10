package main

import (
	"net"

	"github.com/jon20/grpc-stream-sample/server/grpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()

	handler.NewUploadServer(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}

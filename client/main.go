package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jon20/grpc-stream-sample/proto"

	"google.golang.org/grpc"
)

func main() {
	connect, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())

	defer connect.Close()
	uploadhalder := upload.NewUploadHandlerClient(connect)
	strea, err := uploadhalder.Upload(context.Background())
	err = Upload(strea)
	if err != nil {
		fmt.Println(err)
	}
}

func Upload(stream upload.UploadHandler_UploadClient) error {
	file, _ := os.Open("./sample.mp4")
	defer file.Close()
	buf := make([]byte, 1024)

	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.Send(&upload.UploadRequest{VideoData: buf})
	}
	resp, _ := stream.CloseAndRecv()
	aa := resp.GetUploadStatus
	fmt.Println(aa)
	return nil
}

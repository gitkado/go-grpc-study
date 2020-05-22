package main

import (
	"context"
	"log"
	"io"

	"google.golang.org/grpc"

	pb "github.com/gitkado/go-grpc-study/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewAuthorizationClient(conn)

	notes := []*pb.CheckAuthzMessage{
		{UserId: "gitkado"},
		{UserId: "hoge"},
		{},
	}

	stream, err := client.CheckAuthz(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Receive IsCheck %t", in.IsCheck)
		}
	}()

	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
		log.Printf("Send UserId %s", note.UserId)
	}
	stream.CloseSend()
	<-waitc
}

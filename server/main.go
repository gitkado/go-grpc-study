package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/gitkado/go-grpc-study/pb"
	service "github.com/gitkado/go-grpc-study/services"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	AuthorizationService := &service.AuthorizationServerService{}
	pb.RegisterAuthorizationServer(server, AuthorizationService)
	if err := server.Serve(listenPort); err != nil {
		log.Fatalln(err)
	}
}

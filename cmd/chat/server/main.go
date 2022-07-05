package main

import (
	genchat "github.com/manny2014/golang_grpc/internal/gen/chat"
	"github.com/manny2014/golang_grpc/internal/pkg/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	log.Println("starting listener")

	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	log.Println("successfully started listener")

	chatServer := chat.Server{}

	grpcServer := grpc.NewServer()

	genchat.RegisterChatServiceServer(grpcServer, &chatServer)

	log.Println("starting grpc server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve on port 9000: %v", err)
	}
}

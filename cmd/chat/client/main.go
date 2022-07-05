package main

import (
	"context"
	genchat "github.com/manny2014/golang_grpc/internal/gen/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	log.Println("starting client")

	//conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	// New way to do insecure
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	defer conn.Close()

	c := genchat.NewChatServiceClient(conn)

	msg := genchat.Message{
		Body: "Hello from client",
	}

	response, err := c.SayHello(context.Background(), &msg)

	if err != nil {
		log.Fatalf("error sending hello: %v", err)
	}

	log.Printf("response from server: %s", response.Body)
}

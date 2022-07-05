package chat

import (
	"context"
	"github.com/manny2014/golang_grpc/internal/gen/chat"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, msg *chat.Message) (*chat.Message, error) {
	log.Printf("received msg body from client: %s", msg.Body)
	return &chat.Message{
		Body: "Hello from server",
	}, nil
}

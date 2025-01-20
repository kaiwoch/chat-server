package main

import (
	pb "github.com/KamigamiNoGigan/chat-server/pkg/chat_api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedChatApiServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterChatApiServer(s, &server{})

	log.Printf("server listening at %v", port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

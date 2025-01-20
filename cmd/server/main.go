package main

import (
	pb "github.com/KamigamiNoGigan/chat-server/pkg/chat_api_v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"context"
	"math/rand"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedChatApiServer
}

func(s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{Id: rand.Int63n(1000)}, nil
}

func(s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Println(in.Id)
	return &emptypb.Empty{}, nil
}

func(s *server) SendMessage(ctx context.Context, in *pb.SMRequest) (*emptypb.Empty, error) {
	log.Println("from:", in.From)
	log.Println("text:", in.Text)
	log.Println("timestamp:", in.Timestamp)
	return &emptypb.Empty{}, nil
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

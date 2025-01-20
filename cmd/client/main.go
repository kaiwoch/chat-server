package main

import (
	"google.golang.org/grpc"
	"context"
	"time"
	"log"
	pb "github.com/KamigamiNoGigan/chat-server/pkg/chat_api_v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := c.Create(ctx, &pb.CreateRequest{Usernames: []string{"kaiwoch", "dandadan"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)

	_, err = c.Delete(ctx, &pb.DeleteRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.SendMessage(ctx, &pb.SMRequest{From: "kaiwoch", Text: "Hi!", Timestamp: timestamppb.New(time.Now())})
	if err != nil {
		log.Fatal(err)
	}
}
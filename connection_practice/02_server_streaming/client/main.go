package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"connection_practice/pb"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOrderManagementClient(conn)

	// 타임아웃 넉넉하게 설정 (스트리밍이 끝날 때까지 기다려야 하므로)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 1. 서버에 스트리밍 요청 (Stream 객체 획득)
	stream, err := c.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Google Device"})
	if err != nil {
		log.Fatalf("Error calling SearchOrders: %v", err)
	}

	log.Println("Start receiving stream...")

	// 2. 스트림 수신 루프
	for {
		// stream.Recv()는 메시지가 올 때까지 블로킹됨
		order, err := stream.Recv()

		// [핵심] io.EOF는 에러가 아니라 "데이터 전송 끝"을 의미
		if err == io.EOF {
			log.Println("Stream finished (EOF)")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		log.Printf("Received Order: ID=%s, Desc=%s", order.Id, order.Description)
	}
}

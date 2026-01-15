package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"connection_practice/pb"
)

type server struct {
	pb.UnimplementedOrderManagementServer
}

// SearchOrders: Server Streaming 구현
func (s *server) SearchOrders(query *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	log.Printf("Searching for orders with query: %s", query.Value)

	// 시나리오: 검색된 주문이 3개 있다고 가정하고 하나씩 스트리밍
	for i := 0; i < 3; i++ {
		// 1초씩 딜레이를 주어 스트리밍 효과 확인
		time.Sleep(time.Second * 1)

		// 보낼 데이터 생성
		order := &pb.Order{
			Id:          fmt.Sprintf("10%d", i),
			Items:       []string{"Google Pixel", "Chromecast"},
			Description: fmt.Sprintf("Search result %d for %s", i+1, query.Value),
			Price:       float32(100.0 + float32(i)*10),
		}

		// [핵심] stream.Send()로 클라이언트에게 메시지 전송
		if err := stream.Send(order); err != nil {
			log.Printf("Failed to send order: %v", err)
			return err
		}
		log.Printf("Sent Order ID: %s", order.Id)
	}

	// nil을 반환하면 스트림 종료(EOF)를 의미
	log.Println("Finished sending orders.")
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})

	log.Printf("Server Streaming Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/client/ecommerce"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Product examples
	products := []struct {
		name        string
		description string
		price       float32
	}{
		{
			name:        "삼성 갤럭시 S24",
			description: "최신 삼성 갤럭시 S24. 강력한 프로세서와 고화질 디스플레이가 탑재되었습니다.",
			price:       1199.0,
		},
		{
			name:        "LG 올레드 TV 65인치",
			description: "LG OLED TV 65인치. 완벽한 검은색과 생생한 색감으로 최고의 화질을 경험하세요.",
			price:       3500.0,
		},
		{
			name:        "소니 WH-1000XM5 헤드폰",
			description: "소니 프리미엄 노이즈 캔슬링 헤드폰. 최상의 음질과 편안한 착용감을 제공합니다.",
			price:       399.0,
		},
		{
			name:        "애플 아이패드 프로 12.9",
			description: "애플 아이패드 프로 12.9인치. M4 칩 탑재로 뛰어난 성능을 자랑합니다.",
			price:       1299.0,
		},
		{
			name:        "현대 코나 하이브리드",
			description: "현대 신형 코나 하이브리드. 연비 효율성과 세련된 디자인이 결합되었습니다.",
			price:       28500.0,
		},
	}

	// Add products and retrieve them
	for _, product := range products {
		r, err := c.AddProduct(ctx, &pb.Product{
			Name:        product.name,
			Description: product.description,
			Price:       product.price,
		})
		if err != nil {
			log.Fatalf("Could not add product: %v", err)
		}
		log.Printf("상품 ID: %s 추가됨", r.Value)

		// Retrieve the product
		retrieved, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
		if err != nil {
			log.Fatalf("Could not get product: %v", err)
		}
		log.Printf("상품: %v\n", retrieved.String())
	}
}

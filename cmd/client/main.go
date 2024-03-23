package main

import (
	"context"
	"fmt"
	"log"

	"github.com/WalterPaes/go-grpc-crud/config"
	"github.com/WalterPaes/go-grpc-crud/internal/model"
	pb "github.com/WalterPaes/go-grpc-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(cfg.ServerUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	resp, err := client.Create(context.Background(), &pb.ProductRequest{
		Name:        "TV Samsung",
		Category:    "Eletroeletrônicos",
		Description: "TV Samsung 50 Polegadas",
		Price:       1999.99,
	})
	if err != nil {
		log.Fatal(err)
	}

	product, err := client.FindById(context.Background(), &pb.FindProductRequest{
		Id: resp.Id,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", model.Product{
		ID:          uint(product.Id),
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		Price:       float64(product.Price),
	})
}

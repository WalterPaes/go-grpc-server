package main

import (
	"log"
	"net"

	"github.com/WalterPaes/go-grpc-crud/config"
	"github.com/WalterPaes/go-grpc-crud/internal/model"
	"github.com/WalterPaes/go-grpc-crud/internal/repositories"
	"github.com/WalterPaes/go-grpc-crud/internal/services"
	"github.com/WalterPaes/go-grpc-crud/pkg/database"
	pb "github.com/WalterPaes/go-grpc-crud/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDB(sqlite.Open(cfg.DbDSN))
	db.AutoMigrate(&model.Product{})

	listen, err := net.Listen("tcp", cfg.ServerUrl)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	productRepository := repositories.NewProductRepository(db, cfg.DbTimeout)
	productServiceServer := services.NewProductServiceServer(productRepository)

	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, productServiceServer)

	server.Serve(listen)
}

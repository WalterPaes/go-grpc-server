package services

import (
	"context"

	"github.com/WalterPaes/go-grpc-crud/internal/model"
	"github.com/WalterPaes/go-grpc-crud/internal/repositories"
	pb "github.com/WalterPaes/go-grpc-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer
	productRepository repositories.ProductRepository
}

func NewProductServiceServer(productRepository repositories.ProductRepository) *ProductServiceServer {
	return &ProductServiceServer{
		productRepository: productRepository,
	}
}

func (s *ProductServiceServer) Create(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	p, err := s.productRepository.Save(ctx, &model.Product{
		Name:        req.GetName(),
		Category:    req.GetCategory(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()),
	})
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Error: %s", err.Error())
	}

	return &pb.ProductResponse{
		Id:          int32(p.ID),
		Name:        p.Name,
		Category:    p.Category,
		Description: p.Description,
		Price:       float32(p.Price),
	}, nil
}

func (s *ProductServiceServer) FindById(ctx context.Context, req *pb.FindProductRequest) (*pb.ProductResponse, error) {
	p, err := s.productRepository.Find(ctx, int(req.GetId()))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Error: %s", err.Error())
	}

	return &pb.ProductResponse{
		Id:          int32(p.ID),
		Name:        p.Name,
		Category:    p.Category,
		Description: p.Description,
		Price:       float32(p.Price),
	}, nil
}

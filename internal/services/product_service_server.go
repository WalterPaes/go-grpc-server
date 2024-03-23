package services

import (
	"context"
	"log"

	"github.com/WalterPaes/go-grpc-crud/internal/model"
	"github.com/WalterPaes/go-grpc-crud/internal/repositories"
	pb "github.com/WalterPaes/go-grpc-crud/proto"
	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	product := &model.Product{
		Name:        req.GetName(),
		Category:    req.GetCategory(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()),
	}

	validate := validator.New()
	err := validate.Struct(product)
	if err != nil {
		grpcStatus := status.New(codes.InvalidArgument, "bad request")
		badRequestErrDetails := &errdetails.BadRequest{}

		for _, err := range err.(validator.ValidationErrors) {
			v := &errdetails.BadRequest_FieldViolation{
				Field:       err.Field(),
				Description: err.Error(),
			}
			badRequestErrDetails.FieldViolations = append(badRequestErrDetails.FieldViolations, v)
		}

		st, err := grpcStatus.WithDetails(badRequestErrDetails)
		if err != nil {
			log.Fatalf("Unexpected error attaching metadata: %v", err)
		}
		return nil, st.Err()
	}

	p, err := s.productRepository.Save(ctx, product)
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

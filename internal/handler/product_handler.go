package handler

import (
	pb "github.com/WalterPaes/go-grpc-crud/proto"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
}

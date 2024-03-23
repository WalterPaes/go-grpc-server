package repositories

import (
	"context"

	"github.com/WalterPaes/go-grpc-crud/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(*model.Product) (*model.Product, error)
	Find(id int) (*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Save(product *model.Product) (*model.Product, error) {
	result := r.db.WithContext(context.Background()).Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *productRepository) Find(id int) (*model.Product, error) {
	var product model.Product
	result := r.db.WithContext(context.Background()).Find(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

package repositories

import (
	"context"
	"time"

	"github.com/WalterPaes/go-grpc-crud/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(context.Context, *model.Product) (*model.Product, error)
	Find(ctx context.Context, id int) (*model.Product, error)
}

type productRepository struct {
	db        *gorm.DB
	dbTimeout time.Duration
}

func NewProductRepository(db *gorm.DB, dbTimeout time.Duration) *productRepository {
	return &productRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *productRepository) Save(ctx context.Context, product *model.Product) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result := r.db.WithContext(ctx).Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *productRepository) Find(ctx context.Context, id int) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var product model.Product
	result := r.db.WithContext(ctx).Find(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

package repositories

import (
	"database/sql"
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/WalterPaes/go-grpc-crud/internal/model"
	"github.com/WalterPaes/go-grpc-crud/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "test.db"

type RepositorySuite struct {
	suite.Suite
	conn              *sql.DB
	DB                *gorm.DB
	mock              sqlmock.Sqlmock
	productRepository *productRepository
	product           *model.Product
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (rs *RepositorySuite) SetupSuite() {
	var (
		err error
	)

	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)

	rs.DB = database.NewDB(sqlite.Open(dbName))
	rs.DB.AutoMigrate(&model.Product{})

	rs.productRepository = NewProductRepository(rs.DB)
	assert.IsType(rs.T(), &productRepository{}, rs.productRepository)

	rs.product = &model.Product{
		ID:          1,
		Name:        "Notebook Dell",
		Category:    "Informática",
		Description: "Notebook Dell i5",
		Price:       5000.99,
	}
}

func (rs *RepositorySuite) AfterTest(_, _ string) {
	assert.NoError(rs.T(), os.Remove(dbName))
	assert.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *RepositorySuite) Test_productRepository() {
	rs.T().Run("Save", func(t *testing.T) {
		p, err := rs.productRepository.Save(rs.product)
		assert.NoError(rs.T(), err)
		assert.Equal(rs.T(), rs.product, p)
	})

	rs.T().Run("Find", func(t *testing.T) {
		p, err := rs.productRepository.Find(int(rs.product.ID))
		assert.NoError(rs.T(), err)
		assert.Equal(rs.T(), rs.product, p)
	})
}

package infra

import (
	"context"
	"fmt"
	"os"
	"rest-playground/common/postgresql"
	"rest-playground/domain"
	"rest-playground/repository"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var productRepository repository.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M){
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx,postgresql.Config{
		Host:                  "127.0.0.1",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnection:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = repository.NewProductRepository(dbPool)
	fmt.Println("Before all tests")


	exitCode := m.Run()

	fmt.Println("After all tests")
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Computer",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Dishwasher",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Decoration Store",
		},
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Computer",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Dishwasher",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}


func TestAddProduct(t *testing.T) {
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "Monitor",
			Price:    1000.0,
			Discount:0.0,
			Store:    "ABC TECH",
		},
	}
	newProduct := domain.Product{
		Name: "Monitor",
		Price: 1000.0,
		Discount: 0.0,
		Store: "ABC TECH",
	}
	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(newProduct)
		addedProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(addedProducts))
		assert.Equal(t, expectedProducts, addedProducts)
	})

	clear(ctx, dbPool)
}


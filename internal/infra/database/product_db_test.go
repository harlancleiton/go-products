package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/harlancleiton/go-products/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.5)
	assert.NoError(t, err)
	productDb := NewProduct(db)
	err = productDb.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	amountOfProducts := 10
	for i := 0; i < amountOfProducts; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64())
		assert.NoError(t, err)
		db.Create(product)
	}

	productDb := NewProduct(db)
	products, err := productDb.FindAll(1, amountOfProducts, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, amountOfProducts)
	assert.Equal(t, products[0].Name, "Product 0")
	assert.Equal(t, products[amountOfProducts-1].Name, fmt.Sprintf("Product %d", amountOfProducts-1))
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	randomProductName := fmt.Sprintf("Product %d", rand.Intn(100))
	randomProductPrice := rand.Float64()
	product, err := entity.NewProduct(randomProductName, randomProductPrice)
	assert.NoError(t, err)
	db.Create(product)

	productDb := NewProduct(db)
	productFound, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, productFound.Name, randomProductName)
	assert.Equal(t, productFound.Price, randomProductPrice)
}

func TestUpdateProduct(
	t *testing.T,
) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	randomProductName := fmt.Sprintf("Product %d", rand.Intn(100))
	randomProductPrice := rand.Float64()
	product, err := entity.NewProduct(randomProductName, randomProductPrice)
	assert.NoError(t, err)
	db.Create(product)

	productDb := NewProduct(db)
	productFound, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, productFound.Name, randomProductName)
	assert.Equal(t, productFound.Price, randomProductPrice)

	newRandomProductName := fmt.Sprintf("Product %d", rand.Intn(100))
	newRandomProductPrice := rand.Float64()
	productFound.Name = newRandomProductName
	productFound.Price = newRandomProductPrice
	err = productDb.Update(*productFound)
	assert.NoError(t, err)

	productFound, err = productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, productFound.Name, newRandomProductName)
	assert.Equal(t, productFound.Price, newRandomProductPrice)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	randomProductName := fmt.Sprintf("Product %d", rand.Intn(100))
	randomProductPrice := rand.Float64()
	product, err := entity.NewProduct(randomProductName, randomProductPrice)
	assert.NoError(t, err)
	db.Create(product)

	productDb := NewProduct(db)
	err = productDb.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)
}

package entity

import (
	"testing"

	"github.com/harlancleiton/go-products/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	id, _ := entity.NewIDFromString("123e4567-e89b-12d3-a456-426614174000")
	name := "Product 1"
	price := 10.5

	product, err := NewProduct(id, name, price)

	assert.Nil(t, err)
	assert.Equal(t, id, product.ID)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, price, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	id, _ := entity.NewIDFromString("123e4567-e89b-12d3-a456-426614174000")
	name := ""
	price := 10.5

	product, err := NewProduct(id, name, price)

	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	id, _ := entity.NewIDFromString("123e4567-e89b-12d3-a456-426614174000")
	name := "Product 1"
	price := 0.0

	product, err := NewProduct(id, name, price)

	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

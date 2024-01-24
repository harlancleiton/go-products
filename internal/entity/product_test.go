package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := "Product 1"
	price := 10.5

	product, err := NewProduct(name, price)

	assert.Nil(t, err)
	assert.NotNil(t, product.ID)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, price, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	name := ""
	price := 10.5

	product, err := NewProduct(name, price)

	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	name := "Product 1"
	price := 0.0

	product, err := NewProduct(name, price)

	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

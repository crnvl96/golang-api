package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product", 10.99)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product", p.Name)
	assert.Equal(t, 10.99, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	_, err := NewProduct("", 10.99)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	_, err := NewProduct("Product", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsLessThanMinimum(t *testing.T) {
	_, err := NewProduct("Product", -1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product", 10.99)
	assert.Nil(t, err)
	assert.Nil(t, p.Validate())
}

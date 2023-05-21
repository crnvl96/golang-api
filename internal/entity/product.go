package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/crnvl96/go-api/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrNameIsRequired  = errors.New("name is required")
	ErrIdIsInvalid     = errors.New("id is invalid")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("price is invalid")
)

type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIdIsInvalid
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

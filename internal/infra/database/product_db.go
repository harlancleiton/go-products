package database

import (
	"github.com/harlancleiton/go-products/internal/entity"
	"gorm.io/gorm"
)

const DefaultSort = "asc"
const DefaultPage = 1
const DefaultLimit = 10

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product

	if sort != "asc" && sort != "desc" {
		sort = DefaultSort
	}
	if page <= 0 {
		page = DefaultPage
	}
	if limit <= 0 {
		limit = DefaultLimit
	}

	offset := (page - 1) * limit
	err := p.DB.Offset(offset).Limit(limit).Order("created_at " + sort).Find(&products).Error
	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.Where("id = ?", id).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())

	if err != nil {
		return err
	}

	return p.DB.Save(&product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)

	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}

package interfaces

import "depogunabangunan/apps/model"

// IProductRepository defines the interface for product repository operations
type IProductService interface {
	GetProductByID(id int64) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int64) error
}

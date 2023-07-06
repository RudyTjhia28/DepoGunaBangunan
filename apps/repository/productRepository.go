package repository

import (
	"database/sql"
	"depogunabangunan/apps/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProductByID(id int64) (*model.Product, error) {
	product := &model.Product{}
	query := "SELECT id, name, price FROM products WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Product not found
		}
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2)"
	_, err := r.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	query := "UPDATE products SET name = $1, price = $2 WHERE id = $3"
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(id int64) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(name string) ([]models.Product, error) {
	return s.repo.GetProducts(name)
}

func (s *ProductService) CreateProduct(data *models.Product) error {
	return s.repo.CreateProduct(data)
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) UpdateProductByID(product *models.Product) error {
	return s.repo.UpdateProductByID(product)
}
func (s *ProductService) DeleteProductByID(id int) error {
	return s.repo.DeleteProductByID(id)
}

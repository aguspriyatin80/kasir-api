package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories() ([]models.Category, error) {
	return s.repo.GetCategories()
}

func (s *CategoryService) CreateCategory(data *models.Category) error {
	return s.repo.CreateCategory(data)
}

func (s *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) UpdateCategoryByID(Category *models.Category) error {
	return s.repo.UpdateCategoryByID(Category)
}
func (s *CategoryService) DeleteCategoryByID(id int) error {
	return s.repo.DeleteCategoryByID(id)
}

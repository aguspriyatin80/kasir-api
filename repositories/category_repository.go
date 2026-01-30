package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetCategories() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Categorys := make([]models.Category, 0)
	for rows.Next() {
		var p models.Category
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return nil, err
		}
		Categorys = append(Categorys, p)
	}

	return Categorys, nil
}

func (repo *CategoryRepository) CreateCategory(Category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, Category.Name, Category.Description).Scan(&Category.ID)
	return err
}

// GetCategoryByID - ambil kategori by ID
func (repo *CategoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	var p models.Category
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("kategori tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// UpdateCategoryByID - update category by ID
func (repo *CategoryRepository) UpdateCategoryByID(Category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	result, err := repo.db.Exec(query, Category.Name, Category.Description, Category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("kategori tidak ditemukan")
	}

	return nil
}

// DeleteCategoryByID - delete kategori by ID
func (repo *CategoryRepository) DeleteCategoryByID(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("kategori tidak ditemukan")
	}

	return err
}

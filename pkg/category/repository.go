package category

import (
	"database/sql"
	"log"

	"github.com/iugstav/colatech-api/internal/entities"
	"github.com/jmoiron/sqlx"
)

type CategoriesRepository struct {
	DB *sqlx.DB
}

func GenerateNewCategoriesRepository(db *sqlx.DB) entities.ICategoriesRepository {
	return &CategoriesRepository{DB: db}
}

func (r *CategoriesRepository) Create(category *entities.Category) error {
	query := `INSERT INTO categories(id, name) VALUES(:id, :name)`

	_, err := r.DB.NamedExec(query, category)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepository) GetAll() ([]*entities.Category, error) {
	var categories []*entities.Category

	err := r.DB.Select(&categories, "SELECT * FROM categories ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoriesRepository) GetById(id string) (*entities.Category, error) {
	var category entities.Category

	err := r.DB.Get(&category, "SELECT * FROM categories WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoriesRepository) UpdateName(category *entities.Category) error {
	_, err := r.DB.Exec("UPDATE categories SET name=$1 WHERE id=$2", category.Name, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepository) Delete(id string) error {
	_, err := r.DB.Exec(`DELETE FROM categories WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepository) Exists(id string) bool {
	var exists bool

	query := `SELECT EXISTS (SELECT id FROM categories WHERE id=$1)`

	err := r.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatalf("Error while checking existence of category: %v", err.Error())
		}
	}

	return true
}

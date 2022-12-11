package category

import "github.com/jmoiron/sqlx"

type ICategoriesRepository interface {
	Create(category *Category) error
	GetAll() ([]*Category, error)
	GetById(id string) (*Category, error)
	UpdateName(category *Category) error
	Delete(id string) error
}

type CategoriesRepository struct {
	DB *sqlx.DB
}

func GenerateNewCategoriesRepository(db *sqlx.DB) *CategoriesRepository {
	return &CategoriesRepository{DB: db}
}

func (r *CategoriesRepository) Create(category *Category) error {
	query := `INSERT INTO categories(id, name) VALUES(:id, :name)`

	_, err := r.DB.NamedExec(query, category)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepository) GetAll() ([]*Category, error) {
	var categories []*Category

	err := r.DB.Select(&categories, "SELECT * FROM categories ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoriesRepository) GetById(id string) (*Category, error) {
	var category Category

	err := r.DB.Get(&category, "SELECT * FROM categories WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoriesRepository) UpdateName(category *Category) error {
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

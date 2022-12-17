package entities

type Category struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type ICategoriesRepository interface {
	Create(category *Category) error
	GetAll() ([]*Category, error)
	GetById(id string) (*Category, error)
	UpdateName(category *Category) error
	Delete(id string) error
	Exists(id string) bool
}

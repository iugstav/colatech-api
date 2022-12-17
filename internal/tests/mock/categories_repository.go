package mock

import (
	"errors"

	"github.com/iugstav/colatech-api/internal/entities"
)

type CategoriesRepositoryMock struct {
	DB []*entities.Category
}

func GenerateNewMockedCategoriesRepository() *CategoriesRepositoryMock {
	return &CategoriesRepositoryMock{
		DB: []*entities.Category{},
	}
}

func (rm *CategoriesRepositoryMock) Create(category *entities.Category) error {
	rm.DB = append(rm.DB, category)

	return nil
}

func (rm *CategoriesRepositoryMock) GetAll() ([]*entities.Category, error) {
	return rm.DB, nil
}

func (rm *CategoriesRepositoryMock) GetById(id string) (*entities.Category, error) {
	for _, c := range rm.DB {
		if c.ID == id {
			return c, nil
		}
	}

	return nil, errors.New("category does not exists")
}

func (rm *CategoriesRepositoryMock) UpdateName(category *entities.Category) error {
	for _, c := range rm.DB {
		if c.ID == category.ID {
			c.Name = category.Name
		}
	}

	return errors.New("category does not exists")
}

func (rm *CategoriesRepositoryMock) Delete(id string) error {
	var newMock []*entities.Category

	for _, c := range rm.DB {
		if c.ID != id {
			newMock = append(newMock, c)
		}
	}

	rm.DB = newMock

	return errors.New("category does not exists")
}

func (rm *CategoriesRepositoryMock) Exists(id string) bool {
	for _, c := range rm.DB {
		if c.ID == id {
			return true
		}
	}

	return false
}

package category

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type ICategoryService interface {
	Create(name string) (*Category, error)
	GetAll() ([]*Category, error)
	UpdateName(category *Category) error
	Delete(id string) error
}

type CategoryService struct {
	CategoriesRepository ICategoriesRepository
}

func GenerateNewCategoryService(repository ICategoriesRepository) *CategoryService {
	return &CategoryService{CategoriesRepository: repository}
}

func (s *CategoryService) Create(name string) (*Category, error) {
	categoryId := uuid.New().String()

	category := Category{
		ID:   categoryId,
		Name: name,
	}

	if err := s.CategoriesRepository.Create(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) GetAll() ([]*Category, error) {
	response, err := s.CategoriesRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *CategoryService) UpdateName(category *Category) error {
	_, err := uuid.Parse(category.ID)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}

	responseErr := s.CategoriesRepository.UpdateName(category)
	if responseErr != nil {
		return responseErr
	}

	return nil
}

func (s *CategoryService) Delete(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}

	responseErr := s.CategoriesRepository.Delete(id)
	if responseErr != nil {
		return responseErr
	}

	return nil
}

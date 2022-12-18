package category

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/entities"
	"github.com/iugstav/colatech-api/internal/validation"
)

type ICategoryService interface {
	Create(name string) (*entities.Category, error)
	GetAll() ([]*entities.Category, error)
	UpdateName(category *entities.Category) error
	Delete(id string) error
}

type CategoryService struct {
	CategoriesRepository entities.ICategoriesRepository
}

func GenerateNewCategoryService(repository entities.ICategoriesRepository) *CategoryService {
	return &CategoryService{CategoriesRepository: repository}
}

func (s *CategoryService) Create(name string) (*entities.Category, error) {
	categoryId := uuid.New().String()

	if !validation.IsSanitized(name) {
		invalidNameError := fmt.Errorf("invalid name format %s", name)

		return nil, invalidNameError
	}

	category := entities.Category{
		ID:   categoryId,
		Name: name,
	}

	if err := s.CategoriesRepository.Create(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) GetAll() ([]*entities.Category, error) {
	response, err := s.CategoriesRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *CategoryService) UpdateName(category *entities.Category) error {
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

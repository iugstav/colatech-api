package category

import (
	"testing"

	"github.com/iugstav/colatech-api/internal/tests/mock"
)

func TestCreateCategory(t *testing.T) {
	mocked := mock.GenerateNewMockedCategoriesRepository()
	s := GenerateNewCategoryService(mocked)

	s.Create("categoria 1")

	if len(mocked.DB) != 1 {
		t.Errorf("expected db length to be 1, but got 0")
	}
}

func TestShould_Not_Create_Category_With_Invalid_Name(t *testing.T) {
	mocked := mock.GenerateNewMockedCategoriesRepository()
	s := GenerateNewCategoryService(mocked)

	_, err := s.Create("categoria@1")
	if err != nil {
		if err.Error() != "invalid name format categoria@1" {
			t.Errorf("test got an error, but wasn't the expected one.")
		}
	} else {
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	mocked := mock.GenerateNewMockedCategoriesRepository()
	s := GenerateNewCategoryService(mocked)

	s.Create("categoria 1")
	s.Create("categoria 2")

	response, _ := s.GetAll()

	if len(response) != 2 {
		t.Errorf("returned response does not match with quantity of added values")
	}
	if response[0].Name != "categoria 1" || response[1].Name != "categoria 2" {
		t.Errorf("returned response has invalid values inside")
	}
}

func TestUpdateName(t *testing.T) {
	mocked := mock.GenerateNewMockedCategoriesRepository()
	s := GenerateNewCategoryService(mocked)

	category, _ := s.Create("categoria 1")

	category.Name = "novo nome"

	err := s.UpdateName(category)

	if mocked.DB[0].Name != "novo nome" {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	mocked := mock.GenerateNewMockedCategoriesRepository()
	s := GenerateNewCategoryService(mocked)

	category, _ := s.Create("categoria 1")

	s.Delete(category.ID)

	if len(mocked.DB) != 0 {
		t.Errorf("expected db length to be 0, but got %d", len(mocked.DB))
	}
}

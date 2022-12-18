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

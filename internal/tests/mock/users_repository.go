package mock

import (
	"fmt"

	"github.com/iugstav/colatech-api/internal/entities"
)

type UsersRepositoryMock struct {
	DB []*entities.User
}

func GenerateNewMockedUsersRepository() *UsersRepositoryMock {
	return &UsersRepositoryMock{
		DB: []*entities.User{},
	}
}

func (rm *UsersRepositoryMock) Create(user *entities.User) error {
	rm.DB = append(rm.DB, user)

	return nil
}

func (rm *UsersRepositoryMock) GetById(id string) (*entities.User, error) {
	for _, u := range rm.DB {
		if u.ID == id {
			return u, nil
		}
	}

	errorMsg := fmt.Errorf("GetById: user with provided id %s does not exists", id)

	return nil, errorMsg
}

func (rm *UsersRepositoryMock) GetByEmail(email string) (*entities.User, error) {
	for _, u := range rm.DB {
		if u.Email == email {
			return u, nil
		}
	}

	errorMsg := fmt.Errorf("GetById: user with provided email %s does not exists", email)

	return nil, errorMsg
}

func (rm *UsersRepositoryMock) UploadImage(dto *entities.UploadProfileImageToPersistence) error {
	for _, u := range rm.DB {
		if u.ID == dto.ID {
			u.ImageURL = dto.ProfileImageURL
		}
	}

	errorMsg := fmt.Errorf("GetById: user with provided id %s does not exists", dto.ID)

	return errorMsg
}

func (rm *UsersRepositoryMock) Exists(id string) bool {
	for _, u := range rm.DB {
		if u.ID == id {
			return true
		}
	}

	return false
}

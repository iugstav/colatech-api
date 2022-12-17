package mock

import (
	"fmt"

	"github.com/iugstav/colatech-api/pkg/user"
)

type UsersRepositoryMock struct {
	DB []*user.User
}

func GenerateNewMockedUsersRepository() *UsersRepositoryMock {
	return &UsersRepositoryMock{
		DB: []*user.User{},
	}
}

func (rm *UsersRepositoryMock) Create(user *user.User) error {
	rm.DB = append(rm.DB, user)

	return nil
}

func (rm *UsersRepositoryMock) GetById(id string) (*user.User, error) {
	for _, u := range rm.DB {
		if u.ID == id {
			return u, nil
		}
	}

	errorMsg := fmt.Errorf("GetById: user with provided id %s does not exists", id)

	return nil, errorMsg
}

func (rm *UsersRepositoryMock) GetByEmail(email string) (*user.User, error) {
	for _, u := range rm.DB {
		if u.Email == email {
			return u, nil
		}
	}

	errorMsg := fmt.Errorf("GetById: user with provided email %s does not exists", email)

	return nil, errorMsg
}

func (rm *UsersRepositoryMock) UploadImage(dto *user.UploadProfileImageToPersistence) error {
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

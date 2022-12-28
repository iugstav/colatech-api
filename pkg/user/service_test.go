package user

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/entities"
	"github.com/iugstav/colatech-api/internal/tests/mock"
)

func TestCreate(t *testing.T) {
	mocked := mock.GenerateNewMockedUsersRepository()
	s := GenerateNewUserService(mocked)

	userId := uuid.New().String()

	data := entities.CreateUserServiceRequest{
		ID:              userId,
		UserName:        "guzinho",
		FirstName:       "Gustavo",
		LastName:        "Soares",
		Email:           "guguzinho1010@gmail.com",
		Password:        "senhasegura",
		ImageStaticPath: "",
		CreatedAt:       time.Now().Local().Format("2006-01-02 03:04:05"),
	}

	err := s.Create(&data)
	if err != nil {
		t.Error(err)
	}

	if len(mocked.DB) != 1 {
		t.Errorf("expected db length to be 1, but got 0")
	}
}

// TODO: learn how to mock firebase cloud storage

func TestAuthenticate(t *testing.T) {
	mocked := mock.GenerateNewMockedUsersRepository()
	s := GenerateNewUserService(mocked)

	userId := uuid.New().String()

	user := entities.CreateUserServiceRequest{
		ID:              userId,
		UserName:        "guzinho",
		FirstName:       "Gustavo",
		LastName:        "Soares",
		Email:           "guguzinho1010@gmail.com",
		Password:        "senhasegura",
		ImageStaticPath: "",
		CreatedAt:       time.Now().Local().Format("2006-01-02 03:04:05"),
	}

	err := s.Create(&user)
	if err != nil {
		t.Error(err)
	}

	data := entities.AuthenticateUserServiceRequest{
		Email:    user.Email,
		Password: "senhasegura",
	}

	response, err := s.Authenticate(&data)
	if err != nil {
		t.Error(err)
	}

	if response.ID != user.ID {
		t.Errorf("response does not match with user")
	}
}

func TestGetById(t *testing.T) {
	mocked := mock.GenerateNewMockedUsersRepository()
	s := GenerateNewUserService(mocked)

	userId := uuid.New().String()

	user := entities.CreateUserServiceRequest{
		ID:              userId,
		UserName:        "guzinho",
		FirstName:       "Gustavo",
		LastName:        "Soares",
		Email:           "guguzinho1010@gmail.com",
		Password:        "senhasegura",
		ImageStaticPath: "",
		CreatedAt:       time.Now().Local().Format("2006-01-02 03:04:05"),
	}

	err := s.Create(&user)
	if err != nil {
		t.Error(err)
	}

	response, err := s.GetById(userId)
	if err != nil {
		t.Error(err)
	}

	if response.Email != user.Email {
		t.Errorf("response does not match with user")
	}
}

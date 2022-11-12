package user

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/cloud"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(data *CreateUserServiceRequest) error
	Authenticate(data *AuthenticateUserServiceRequest) (*AuthenticateUserResponse, error)
	GetById(id string) (*GetUserByIdResponse, error)
	UploadIMage(data *UploadProfileImageRequest) error
}

type UserService struct {
	UsersRepository IUsersRepository
}

func GenerateNewUserService(repository IUsersRepository) *UserService {
	return &UserService{UsersRepository: repository}
}

func (s *UserService) Create(data *CreateUserServiceRequest) error {
	formattedCreationDate, parseErr := time.Parse("2006-01-02 03:04:05", data.CreatedAt)
	if parseErr != nil {
		return parseErr
	}

	hashedPassword, hashingErr := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if hashingErr != nil {
		errMessage := errors.New("could not hash the given passowrd")

		return errMessage
	}

	user := User{
		ID:        data.ID,
		UserName:  data.UserName,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  string(hashedPassword),
		ImageURL:  "",
		CreatedAt: formattedCreationDate,
	}

	repositoryErr := s.UsersRepository.Create(&user)
	if repositoryErr != nil {
		return repositoryErr
	}

	return nil
}

func (s *UserService) UploadIMage(data *UploadProfileImageRequest) error {
	defer data.File.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bucket, err := cloud.CloudStorage().DefaultBucket()
	if err != nil {
		return nil
	}

	writer := bucket.Object(data.NameToUpload).NewWriter(ctx)
	writer.ContentType = "image/webp"
	writer.Metadata = map[string]string{
		"created-at": time.Now().Format("2016-06-27"),
	}
	defer writer.Close()

	fbytes, err := io.ReadAll(data.File)
	if err != nil {
		return nil
	}

	buf := bytes.NewBuffer(fbytes)

	if _, err = io.Copy(writer, buf); err != nil {
		return nil
	}

	imageURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media",
		os.Getenv("FIREBASE_STORAGE_BUCKET_NAME"),
		"users%2F"+data.NameToUpload)

	d := &UploadProfileImageToPersistence{
		ID:              data.ID,
		ProfileImageURL: imageURL,
	}

	responseErr := s.UsersRepository.UploadImage(d)
	if responseErr != nil {
		return responseErr
	}

	return nil
}

func (s *UserService) Authenticate(data *AuthenticateUserServiceRequest) (*AuthenticateUserResponse, error) {
	var role string

	user, err := s.UsersRepository.GetByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, err
	}

	if user.UserName == "iugstav" {
		role = "author"
	} else {
		role = "reader"
	}

	response := &AuthenticateUserResponse{
		ID:   user.ID,
		Role: role,
	}

	return response, nil
}

func (s *UserService) GetById(id string) (*GetUserByIdResponse, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return nil, errors.New(errMessage)
	}

	response, err := s.UsersRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	user := GetUserByIdResponse{
		UserName:  response.UserName,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
	}

	return &user, nil
}

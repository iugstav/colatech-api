package post

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/cloud"
	"github.com/iugstav/colatech-api/internal/entities"
	"github.com/pkg/errors"
)

type IPostService interface {
	Create(data *entities.CreatePostServiceRequest) (*entities.Post, error)
	GetAll() ([]*entities.Post, error)
	GetAllMinified() ([]*entities.ResumedPost, error)
	GetById(id string) (*entities.GetPostByIdServiceResponse, error)
	UpdateContent(data *entities.UpdatePostContentDTO) error
	UploadImage(data *entities.UploadPostCoverImageRequest) error
	LikePost(dto *entities.LikePostDTO) error
	Delete(id string) error
}

type PostService struct {
	PostsRepository    entities.IPostsRepository
	CommentsRepository entities.ICommentsRepository
}

func GenerateNewPostService(postsRepo entities.IPostsRepository, commentsRepo entities.ICommentsRepository) *PostService {
	return &PostService{PostsRepository: postsRepo, CommentsRepository: commentsRepo}
}

func (s *PostService) Create(data *entities.CreatePostServiceRequest) (*entities.Post, error) {
	formattedPostDate, parseErr := time.Parse("2006-01-02 03:04:05", data.PublishedAt)
	if parseErr != nil {
		return nil, parseErr
	}

	post := &entities.Post{
		ID:            data.ID,
		Title:         data.Title,
		Slug:          data.Slug,
		Intro:         data.Intro,
		CoverImageURL: "",
		Content:       data.Content,
		CategoryID:    data.CategoryID,
		PublishedAt:   formattedPostDate,
	}

	repositoryErr := s.PostsRepository.Create(post)
	if repositoryErr != nil {
		return nil, repositoryErr
	}

	return post, nil
}

func (s *PostService) UploadImage(data *entities.UploadPostCoverImageRequest) error {
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
		"posts%2F"+data.NameToUpload)

	d := &entities.UploadPostCoverImageInPersistence{
		ID:            data.ID,
		CoverImageURL: imageURL,
	}

	responseErr := s.PostsRepository.UploadImage(d)
	if responseErr != nil {
		return responseErr
	}

	return nil
}

func (s *PostService) GetAll() ([]*entities.Post, error) {
	response, err := s.PostsRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostService) GetAllMinified() ([]*entities.ResumedPost, error) {
	response, err := s.PostsRepository.GetAllMinified()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostService) GetById(id string) (*entities.GetPostByIdServiceResponse, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return nil, errors.New(errMessage)
	}

	response, repositoryErr := s.PostsRepository.GetById(id)
	if repositoryErr != nil {
		return nil, err
	}

	cmt, commentErr := s.CommentsRepository.GetAllFromAPost(id)
	if commentErr != nil {
		return nil, err
	}

	post := &entities.GetPostByIdServiceResponse{
		ID:            response.ID,
		Title:         response.Title,
		Slug:          response.Slug,
		Content:       response.Content,
		CoverImageURL: response.CoverImageURL,
		Category: entities.Category{
			ID:   response.CategoryID,
			Name: response.CategoryName,
		},
		PublishedAt: response.PublishedAt,
		Comments:    *cmt,
	}

	return post, nil
}

func (s *PostService) UpdateContent(data *entities.UpdatePostContentDTO) error {
	_, err := uuid.Parse(data.ID)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}

	repositoryErr := s.PostsRepository.UpdateContent(data)
	if repositoryErr != nil {
		return repositoryErr
	}

	return nil
}

func (s *PostService) LikePost(dto *entities.LikePostDTO) error {
	_, err := uuid.Parse(dto.UserID)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}
	_, err = uuid.Parse(dto.PostID)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}

	postExists, userExists := s.PostsRepository.BothUserAndPostExists(dto.UserID, dto.UserID)

	if !postExists && userExists {
		errorMsg := errors.New("this id does not refers to any post")

		return errorMsg
	} else if postExists && !userExists {
		errorMsg := errors.New("this id does not refer to any user")

		return errorMsg
	}

	likeID := uuid.NewString()

	data := &entities.LikePostInPersistence{
		ID:     likeID,
		UserID: dto.UserID,
		PostID: dto.PostID,
	}

	err = s.PostsRepository.LikePost(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) Delete(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}

	repositoryErr := s.PostsRepository.Delete(id)
	if repositoryErr != nil {
		return repositoryErr
	}

	return nil
}

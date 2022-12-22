package comment

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/entities"
)

type ICommentService interface {
	Create(data *entities.CreateCommentServiceRequest) (*entities.Comment, error)
	GetAllFromAPost(postId string) ([]*entities.GetAllFromAPostServiceResponse, error)
	UpdateContent(dto *entities.Comment) error
	Delete(id string) error
}

type CommentService struct {
	CommentsRepository entities.ICommentsRepository
}

func GenerateNewCommentService(commentsRepository entities.ICommentsRepository) *CommentService {
	return &CommentService{CommentsRepository: commentsRepository}
}

func (s *CommentService) Create(data *entities.CreateCommentServiceRequest) (*entities.Comment, error) {
	commentId := uuid.New().String()
	formattedCommentDate, parseErr := time.Parse("2006-01-02 03:04:05", data.PublishedAt)
	if parseErr != nil {
		return nil, parseErr
	}

	var isParentCommentIdValid bool

	if len(data.ParentCommentId) == 0 {
		isParentCommentIdValid = false
	} else {
		isParentCommentIdValid = true
	}

	comment := &entities.Comment{
		ID:       commentId,
		ReaderId: data.ReaderId,
		PostId:   data.PostId,
		ParentCommentId: sql.NullString{
			String: data.ParentCommentId,
			Valid:  isParentCommentIdValid,
		},
		Content:     data.Content,
		PublishedAt: formattedCommentDate,
	}

	if err := s.CommentsRepository.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentService) GetAllFromAPost(postId string) ([]*entities.GetAllFromAPostServiceResponse, error) {
	var comments []*entities.GetAllFromAPostServiceResponse

	_, err := uuid.Parse(postId)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return nil, errors.New(errMessage)
	}

	response, responseErr := s.CommentsRepository.GetAllFromAPost(postId)
	if responseErr != nil {
		return nil, responseErr
	}

	for _, comment := range *response {
		var parentCommentId string

		if comment.ParentCommentId.Valid {
			parentCommentId = comment.ParentCommentId.String
		} else {
			parentCommentId = ""
		}

		dataToAppend := &entities.GetAllFromAPostServiceResponse{
			ID:              comment.ID,
			PostId:          comment.PostId,
			ParentCommentId: parentCommentId,
			Content:         comment.Content,
			PublishedAt:     comment.PublishedAt,
			Reader: entities.ReaderInfoInsideComment{
				ID:        comment.ReaderId,
				FirstName: comment.ReaderFirstName,
				LastName:  comment.ReaderLastName,
			},
		}

		comments = append(comments, dataToAppend)
	}

	return comments, nil
}

func (s *CommentService) UpdateContent(dto *entities.UpdateCommentContentDTO) error {
	_, err := uuid.Parse(dto.ID)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}

	responseErr := s.CommentsRepository.UpdateContent(dto)
	if responseErr != nil {
		return responseErr
	}

	return nil
}

func (s *CommentService) Delete(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return errors.New(errMessage)
	}

	responseErr := s.CommentsRepository.Delete(id)
	if responseErr != nil {
		return responseErr
	}

	return nil
}

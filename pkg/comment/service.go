package comment

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ICommentService interface {
	Create(data *CreateCommentServiceRequest) (*Comment, error)
	GetAllFromAPost(postId string) ([]*GetAllFromAPostServiceResponse, error)
	UpdateContent(dto *UpdateCommentContentDTO) error
	Delete(id string) error
}

type CommentService struct {
	CommentsRepository ICommentsRepository
}

func GenerateNewCommentService(commentsRepository *CommentsRepository) *CommentService {
	return &CommentService{CommentsRepository: commentsRepository}
}

func (s *CommentService) Create(data *CreateCommentServiceRequest) (*Comment, error) {
	commentId := uuid.New().String()
	formattedCommentDate, parseErr := time.Parse("2006-01-02 03:04:05", data.PublishedAt)
	if parseErr != nil {
		return nil, parseErr
	}

	comment := &Comment{
		ID:              commentId,
		ReaderId:        data.ReaderId,
		PostId:          data.PostId,
		ParentCommentId: data.ParentCommentId,
		Content:         data.Content,
		PublishedAt:     formattedCommentDate,
	}

	if err := s.CommentsRepository.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentService) GetAllFromAPost(postId string) ([]*GetAllFromAPostServiceResponse, error) {
	var comments []*GetAllFromAPostServiceResponse

	_, err := uuid.Parse(postId)
	if err != nil {
		errMessage := fmt.Sprintf("Invalid uuid: %v", err.Error())

		return nil, errors.New(errMessage)
	}

	response, responseErr := s.CommentsRepository.GetAllFromAPost(postId)
	if responseErr != nil {
		return nil, responseErr
	}

	for _, comment := range response {
		dataToAppend := &GetAllFromAPostServiceResponse{
			ID:              comment.ID,
			PostId:          comment.PostId,
			ParentCommentId: comment.ParentCommentId,
			Content:         comment.Content,
			PublishedAt:     comment.PublishedAt,
			Reader: ReaderInfoInsideComment{
				ID:        comment.ReaderId,
				FirstName: comment.ReaderFirstName,
				LastName:  comment.ReaderLastName,
			},
		}

		comments = append(comments, dataToAppend)
	}

	return comments, nil
}

func (s *CommentService) UpdateContent(dto *UpdateCommentContentDTO) error {
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

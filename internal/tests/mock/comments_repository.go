package mock

import (
	"fmt"

	"github.com/iugstav/colatech-api/internal/entities"
)

type CommentsRepositoryMock struct {
	DB              []*entities.Comment
	UsersRepository entities.IUsersRepository
}

func GenerateNewMockedCommentsRepository(usersRepo entities.IUsersRepository) *CommentsRepositoryMock {
	return &CommentsRepositoryMock{
		DB:              []*entities.Comment{},
		UsersRepository: usersRepo,
	}
}

func (rm *CommentsRepositoryMock) Create(comment *entities.Comment) error {
	rm.DB = append(rm.DB, comment)

	return nil
}

func (rm *CommentsRepositoryMock) GetAllFromAPost(postId string) (*[]entities.CommentFromPersistence, error) {
	var comments []entities.CommentFromPersistence

	for _, c := range rm.DB {
		if c.PostId == postId {
			user, err := rm.UsersRepository.GetById(c.ReaderId)
			if err != nil {
				unknownReaderError := fmt.Errorf("unknown reader with id %s", c.ReaderId)

				return nil, unknownReaderError
			}

			comment := entities.CommentFromPersistence{
				ID:              c.ID,
				ReaderId:        c.ReaderId,
				ReaderFirstName: user.FirstName,
				ReaderLastName:  user.LastName,
				PostId:          c.PostId,
				ParentCommentId: c.ParentCommentId,
				Content:         c.Content,
				PublishedAt:     c.PublishedAt,
			}

			comments = append(comments, comment)
		}
	}

	return &comments, nil
}

func (rm *CommentsRepositoryMock) UpdateContent(dto *entities.Comment) error {
	for _, c := range rm.DB {
		if c.ID == dto.ID {
			c.Content = dto.Content
		}
	}

	return nil
}

func (rm *CommentsRepositoryMock) Delete(id string) error {
	var newMock []*entities.Comment

	for _, c := range rm.DB {
		if c.ID != id {
			newMock = append(newMock, c)
		}
	}

	rm.DB = newMock

	return nil
}

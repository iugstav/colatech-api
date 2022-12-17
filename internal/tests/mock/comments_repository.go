package mock

import (
	"fmt"

	"github.com/iugstav/colatech-api/pkg/comment"
	"github.com/iugstav/colatech-api/pkg/user"
)

type CommentsRepositoryMock struct {
	DB              []*comment.Comment
	UsersRepository user.IUsersRepository
}

func GenerateNewMockedCommentsRepository(usersRepo user.IUsersRepository) *CommentsRepositoryMock {
	return &CommentsRepositoryMock{
		DB:              []*comment.Comment{},
		UsersRepository: usersRepo,
	}
}

func (rm *CommentsRepositoryMock) Create(comment *comment.Comment) error {
	rm.DB = append(rm.DB, comment)

	return nil
}

func (rm *CommentsRepositoryMock) GetAllFromAPost(postId string) (*[]comment.CommentFromPersistence, error) {
	var comments []comment.CommentFromPersistence

	for _, c := range rm.DB {
		if c.PostId == postId {
			user, err := rm.UsersRepository.GetById(c.ReaderId)
			if err != nil {
				unknownReaderError := fmt.Errorf("unknown reader with id %s", c.ReaderId)

				return nil, unknownReaderError
			}

			comment := comment.CommentFromPersistence{
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

func (rm *CommentsRepositoryMock) UpdateContent(dto *comment.UpdateCommentContentDTO) error {
	for _, c := range rm.DB {
		if c.ID == dto.ID {
			c.Content = dto.NewContent
		}
	}

	return nil
}

func (rm *CommentsRepositoryMock) Delete(id string) error {
	var newMock []*comment.Comment

	for _, c := range rm.DB {
		if c.ID != id {
			newMock = append(newMock, c)
		}
	}

	rm.DB = newMock

	return nil
}

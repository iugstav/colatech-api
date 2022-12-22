package comment

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/entities"
	"github.com/iugstav/colatech-api/internal/tests/mock"
)

func TestCreateComment(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	mocked := mock.GenerateNewMockedCommentsRepository(ur)
	s := GenerateNewCommentService(mocked)

	user := entities.User{
		ID:        "user123",
		UserName:  "guzinho",
		FirstName: "Gustavo",
		LastName:  "Soares",
		Email:     "guguzinho1010@gmail.com",
		Password:  "senhasegura",
		ImageURL:  "",
		CreatedAt: time.Now(),
	}

	ur.Create(&user)

	data := entities.CreateCommentServiceRequest{
		ReaderId:        user.ID,
		PostId:          "post123",
		ParentCommentId: "",
		Content:         "message for random comment",
		PublishedAt:     time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	_, err := s.Create(&data)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllFromAPost(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	mocked := mock.GenerateNewMockedCommentsRepository(ur)
	s := GenerateNewCommentService(mocked)

	postId := uuid.New().String()
	userId := uuid.New().String()

	user := entities.User{
		ID:        userId,
		UserName:  "guzinho",
		FirstName: "Gustavo",
		LastName:  "Soares",
		Email:     "guguzinho1010@gmail.com",
		Password:  "senhasegura",
		ImageURL:  "",
		CreatedAt: time.Now(),
	}

	ur.Create(&user)

	data := []entities.CreateCommentServiceRequest{
		{
			ReaderId:        userId,
			PostId:          postId,
			ParentCommentId: "",
			Content:         "message for random comment",
			PublishedAt:     time.Now().UTC().Format("2006-01-02 03:04:05"),
		},
		{
			ReaderId:        userId,
			PostId:          postId,
			ParentCommentId: "",
			Content:         "message for random comment",
			PublishedAt:     time.Now().UTC().Format("2006-01-02 03:04:05"),
		},
	}

	for _, c := range data {
		_, err := s.Create(&c)
		if err != nil {
			t.Error(err)
		}
	}

	response, err := s.GetAllFromAPost(postId)

	if len(response) == 0 {
		t.Error(err)
	}
}

func TestUpdateContent(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	mocked := mock.GenerateNewMockedCommentsRepository(ur)
	s := GenerateNewCommentService(mocked)

	data := entities.CreateCommentServiceRequest{
		ReaderId:        "user1",
		PostId:          "post123",
		ParentCommentId: "",
		Content:         "message for random comment",
		PublishedAt:     time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	comm, _ := s.Create(&data)

	err := s.UpdateContent(&entities.UpdateCommentContentDTO{
		ID:      comm.ID,
		Content: "new message",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	mocked := mock.GenerateNewMockedCommentsRepository(ur)
	s := GenerateNewCommentService(mocked)

	data := entities.CreateCommentServiceRequest{
		ReaderId:        "user1",
		PostId:          "post123",
		ParentCommentId: "",
		Content:         "message for random comment",
		PublishedAt:     time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	comm, _ := s.Create(&data)

	err := s.Delete(comm.ID)
	if err != nil {
		t.Error(err)
	}

	if len(mocked.DB) > 0 {
		t.Errorf("expected db length 0, but got 1")
	}
}

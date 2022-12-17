package mock

import (
	"github.com/iugstav/colatech-api/internal/entities"
)

type LikesRepositoryMock struct {
	DB []*entities.LikePostInPersistence
}

func GenerateNewMockedLikesRepository() *LikesRepositoryMock {
	return &LikesRepositoryMock{
		DB: []*entities.LikePostInPersistence{},
	}
}

func (rm *LikesRepositoryMock) LikePost(data *entities.LikePostInPersistence) error {
	rm.DB = append(rm.DB, data)

	return nil
}

func (rm *LikesRepositoryMock) DislikePost(data *entities.LikePostInPersistence) error {
	var newMock []*entities.LikePostInPersistence

	for _, l := range rm.DB {
		if l.ID != data.ID {
			newMock = append(newMock, l)
		}
	}

	rm.DB = newMock

	return nil
}

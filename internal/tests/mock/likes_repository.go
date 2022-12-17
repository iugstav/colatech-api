package mock

import "github.com/iugstav/colatech-api/pkg/likes"

type LikesRepositoryMock struct {
	DB []*likes.LikePostInPersistence
}

func GenerateNewMockedLikesRepository() *LikesRepositoryMock {
	return &LikesRepositoryMock{
		DB: []*likes.LikePostInPersistence{},
	}
}

func (rm *LikesRepositoryMock) LikePost(data *likes.LikePostInPersistence) error {
	rm.DB = append(rm.DB, data)

	return nil
}

func (rm *LikesRepositoryMock) DislikePost(data *likes.LikePostInPersistence) error {
	var newMock []*likes.LikePostInPersistence

	for _, l := range rm.DB {
		if l.ID != data.ID {
			newMock = append(newMock, l)
		}
	}

	rm.DB = newMock

	return nil
}

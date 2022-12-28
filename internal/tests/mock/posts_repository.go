package mock

import (
	"fmt"

	"github.com/iugstav/colatech-api/internal/entities"
)

type PostsRepositoryMock struct {
	DB                   []*entities.Post
	CategoriesRepository entities.ICategoriesRepository
	LikesRepository      entities.ILikesRepository
	UsersRepository      entities.IUsersRepository
}

func GenerateNewMockedPostsRepository(categoriesRepo entities.ICategoriesRepository, likesRepo entities.ILikesRepository, usersRepo entities.IUsersRepository) *PostsRepositoryMock {
	return &PostsRepositoryMock{
		DB:                   []*entities.Post{},
		CategoriesRepository: categoriesRepo,
		LikesRepository:      likesRepo,
		UsersRepository:      usersRepo,
	}
}

func (rm *PostsRepositoryMock) Create(p *entities.Post) error {
	rm.DB = append(rm.DB, p)

	return nil
}

func (rm PostsRepositoryMock) GetAll() ([]*entities.Post, error) {
	response := rm.DB

	return response, nil
}

func (rm *PostsRepositoryMock) GetAllMinified() ([]*entities.ResumedPost, error) {
	response := []*entities.ResumedPost{}

	for _, p := range rm.DB {
		data := entities.ResumedPost{
			ID:            p.ID,
			Title:         p.Title,
			Slug:          p.Slug,
			Intro:         p.Intro,
			CategoryID:    p.CategoryID,
			CoverImageURL: p.CoverImageURL,
			PublishedAt:   p.PublishedAt,
		}
		response = append(response, &data)
	}

	return response, nil
}

func (rm *PostsRepositoryMock) GetById(id string) (*entities.GetPostByIdFromRepository, error) {
	for _, p := range rm.DB {
		if p.ID == id {
			ct, err := rm.CategoriesRepository.GetById(p.CategoryID)
			if err != nil {
				return nil, err
			}

			fromPersistence := entities.GetPostByIdFromRepository{
				ID:            p.ID,
				Title:         p.Title,
				Slug:          p.Slug,
				Content:       p.Content,
				CategoryID:    p.CategoryID,
				CategoryName:  ct.Name,
				CoverImageURL: p.CoverImageURL,
				PublishedAt:   p.PublishedAt,
			}

			return &fromPersistence, nil
		}
	}

	errorMsg := fmt.Errorf("GetById: post with provided id %s does not exists", id)

	return nil, errorMsg
}

func (rm *PostsRepositoryMock) UpdateContent(data *entities.UpdatePostContentDTO) error {
	if rm.Exists(data.ID) {
		for _, c := range rm.DB {
			if c.ID == data.ID {
				c.Content = data.NewContent
				return nil
			}
		}
	}

	return nil
}

func (rm *PostsRepositoryMock) UploadImage(data *entities.UploadPostCoverImageInPersistence) error {
	post, err := rm.GetById(data.ID)
	if err != nil {
		return err
	}

	post.CoverImageURL = data.CoverImageURL

	return nil
}

func (rm *PostsRepositoryMock) LikePost(data *entities.LikePostInPersistence) error {
	err := rm.LikesRepository.LikePost(data)
	if err != nil {
		return err
	}

	return nil
}

func (rm *PostsRepositoryMock) Delete(id string) error {
	var newMock []*entities.Post

	for _, p := range rm.DB {
		if p.ID != id {
			newMock = append(newMock, p)
		}
	}

	rm.DB = newMock

	return nil
}

func (rm *PostsRepositoryMock) Exists(id string) bool {
	for _, p := range rm.DB {
		if p.ID == id {
			return true
		}
	}

	return false
}

func (rm *PostsRepositoryMock) BothUserAndPostExists(userId string, postId string) (bool, bool) {
	var userExists bool
	var postExists bool

	if rm.UsersRepository.Exists(userId) {
		userExists = true
	} else {
		userExists = false
	}

	if rm.Exists(postId) {
		postExists = true
	} else {
		postExists = false
	}

	return postExists, userExists
}

package mock

import (
	"fmt"

	"github.com/iugstav/colatech-api/pkg/category"
	"github.com/iugstav/colatech-api/pkg/likes"
	"github.com/iugstav/colatech-api/pkg/post"
	"github.com/iugstav/colatech-api/pkg/user"
)

type PostsRepositoryMock struct {
	DB                   []*post.Post
	CategoriesRepository category.ICategoriesRepository
	LikesRepository      likes.ILikesRepository
	UsersRepository      user.IUsersRepository
}

func GenerateNewMockedPostsRepository(categoriesRepo category.ICategoriesRepository, likesRepo likes.ILikesRepository, usersRepo user.IUsersRepository) *PostsRepositoryMock {
	return &PostsRepositoryMock{
		DB:                   []*post.Post{},
		CategoriesRepository: categoriesRepo,
		LikesRepository:      likesRepo,
		UsersRepository:      usersRepo,
	}
}

func (rm *PostsRepositoryMock) Create(p *post.Post) error {
	rm.DB = append(rm.DB, p)

	return nil
}

func (rm PostsRepositoryMock) GetAll() ([]*post.Post, error) {
	response := rm.DB

	return response, nil
}

func (rm *PostsRepositoryMock) GetAllMinified() ([]*post.ResumedPost, error) {
	response := []*post.ResumedPost{}

	for _, p := range rm.DB {
		data := post.ResumedPost{
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

func (rm *PostsRepositoryMock) GetById(id string) (*post.GetPostByIdFromRepository, error) {
	ct, err := rm.CategoriesRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	for _, p := range rm.DB {
		if p.ID == id {
			fromPersistence := post.GetPostByIdFromRepository{
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

func (rm *PostsRepositoryMock) UpdateContent(data *post.UpdatePostContentDTO) error {
	post, err := rm.GetById(data.ID)
	if err != nil {
		return err
	}

	post.Content = data.NewContent

	return nil
}

func (rm *PostsRepositoryMock) UploadImage(data *post.UploadPostCoverImageInPersistence) error {
	post, err := rm.GetById(data.ID)
	if err != nil {
		return err
	}

	post.CoverImageURL = data.CoverImageURL

	return nil
}

func (rm *PostsRepositoryMock) LikePost(data *likes.LikePostInPersistence) error {
	err := rm.LikesRepository.LikePost(data)
	if err != nil {
		return err
	}

	return nil
}

func (rm *PostsRepositoryMock) Delete(id string) error {
	var newMock []*post.Post

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

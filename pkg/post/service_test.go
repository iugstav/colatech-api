package post

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/entities"
	"github.com/iugstav/colatech-api/internal/tests/mock"
)

func TestCreatePost(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	lr := mock.GenerateNewMockedLikesRepository()
	cr := mock.GenerateNewMockedCategoriesRepository()
	commRep := mock.GenerateNewMockedCommentsRepository(ur)
	mocked := mock.GenerateNewMockedPostsRepository(cr, lr, ur)

	s := GenerateNewPostService(mocked, commRep)

	postId := uuid.New().String()
	categoryId := uuid.New().String()

	cat := entities.Category{
		ID:   categoryId,
		Name: "categoria 1",
	}

	cr.Create(&cat)

	data := entities.CreatePostServiceRequest{
		ID:          postId,
		Title:       "o advento do código",
		Slug:        "code-advent",
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:  categoryId,
		PublishedAt: time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	response, err := s.Create(&data)
	if err != nil {
		t.Error(err)
	}

	if len(mocked.DB) != 1 || response.ID != postId {
		t.Error("expected db length to be 1 but got 0")
	}
}

// TODO: learn how to test firebase cloud storage

func TestGetAll(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	lr := mock.GenerateNewMockedLikesRepository()
	cr := mock.GenerateNewMockedCategoriesRepository()
	commRep := mock.GenerateNewMockedCommentsRepository(ur)
	mocked := mock.GenerateNewMockedPostsRepository(cr, lr, ur)

	s := GenerateNewPostService(mocked, commRep)

	post1 := entities.Post{
		ID:            "post123",
		Title:         "o advento do código",
		Slug:          "code-advent",
		Content:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:    "categoria1",
		PublishedAt:   time.Now().UTC(),
		CoverImageURL: "",
	}
	post2 := entities.Post{
		ID:            "blablabla",
		Title:         "o advento do código",
		Slug:          "code-advent",
		Content:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:    "categoria",
		PublishedAt:   time.Now().UTC(),
		CoverImageURL: "",
	}

	mocked.Create(&post1)
	mocked.Create(&post2)

	response, err := s.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(response) != 2 {
		t.Errorf("expected db length of 2, but got %d", len(response))
	}

	if response[0].ID != "post123" {
		t.Errorf("invalid response data")
	}
}

func TestGetAllMinified(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	lr := mock.GenerateNewMockedLikesRepository()
	cr := mock.GenerateNewMockedCategoriesRepository()
	commRep := mock.GenerateNewMockedCommentsRepository(ur)
	mocked := mock.GenerateNewMockedPostsRepository(cr, lr, ur)

	s := GenerateNewPostService(mocked, commRep)

	post1 := entities.Post{
		ID:            "post123",
		Title:         "o advento do código",
		Slug:          "code-advent",
		Content:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:    "categoria1",
		PublishedAt:   time.Now().UTC(),
		CoverImageURL: "",
	}
	post2 := entities.Post{
		ID:            "blablabla",
		Title:         "o advento do código",
		Slug:          "code-advent",
		Content:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:    "categoria",
		PublishedAt:   time.Now().UTC(),
		CoverImageURL: "",
	}

	mocked.Create(&post1)
	mocked.Create(&post2)

	response, err := s.GetAllMinified()
	if err != nil {
		t.Error(err)
	}

	if len(response) != 2 {
		t.Errorf("expected db length of 2, but got %d", len(response))
	}

	if response[0].ID != "post123" {
		t.Errorf("invalid response data")
	}
}

func TestGetById(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	lr := mock.GenerateNewMockedLikesRepository()
	cr := mock.GenerateNewMockedCategoriesRepository()
	commRep := mock.GenerateNewMockedCommentsRepository(ur)
	mocked := mock.GenerateNewMockedPostsRepository(cr, lr, ur)

	s := GenerateNewPostService(mocked, commRep)

	postId1 := uuid.New().String()
	categoryId := uuid.New().String()

	cat := entities.Category{
		ID:   categoryId,
		Name: "categoria 1",
	}

	mocked.CategoriesRepository.Create(&cat)

	post1 := entities.CreatePostServiceRequest{
		ID:          postId1,
		Title:       "o advento do código",
		Slug:        "code-advent",
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:  categoryId,
		PublishedAt: time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	data, err := s.Create(&post1)
	if err != nil {
		t.Error(err)
	}

	response, err := s.GetById(data.ID)
	if err != nil {
		t.Error(err)
	}

	if response == nil {
		t.Errorf("invalid response from getById")
	}
}

func TestUpdateContent(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	lr := mock.GenerateNewMockedLikesRepository()
	cr := mock.GenerateNewMockedCategoriesRepository()
	commRep := mock.GenerateNewMockedCommentsRepository(ur)
	mocked := mock.GenerateNewMockedPostsRepository(cr, lr, ur)

	s := GenerateNewPostService(mocked, commRep)

	postId1 := uuid.New().String()
	categoryId := uuid.New().String()

	cat := entities.Category{
		ID:   categoryId,
		Name: "categoria 1",
	}

	mocked.CategoriesRepository.Create(&cat)

	post1 := entities.CreatePostServiceRequest{
		ID:          postId1,
		Title:       "o advento do código",
		Slug:        "code-advent",
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:  categoryId,
		PublishedAt: time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	s.Create(&post1)

	newContent := "novo conteudo qualquer coisa blablablablabla"
	data := entities.UpdatePostContentDTO{
		ID:         postId1,
		NewContent: newContent,
	}

	err := s.UpdateContent(&data)
	if err != nil {
		t.Error(err)
	}
}

func TestLikePost(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	lr := mock.GenerateNewMockedLikesRepository()
	cr := mock.GenerateNewMockedCategoriesRepository()
	commRep := mock.GenerateNewMockedCommentsRepository(ur)
	mocked := mock.GenerateNewMockedPostsRepository(cr, lr, ur)

	s := GenerateNewPostService(mocked, commRep)

	postId1 := uuid.New().String()
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

	post1 := entities.CreatePostServiceRequest{
		ID:          postId1,
		Title:       "o advento do código",
		Slug:        "code-advent",
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:  "categoria 1",
		PublishedAt: time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	s.Create(&post1)

	data := entities.LikePostDTO{
		UserID: user.ID,
		PostID: post1.ID,
	}

	err := s.LikePost(&data)
	if err != nil {
		t.Error(err)
	}

	if len(lr.DB) != 1 {
		t.Errorf("expected 1 like in database, but got 0")
	}
}

func TestDelete(t *testing.T) {
	ur := mock.GenerateNewMockedUsersRepository()
	lr := mock.GenerateNewMockedLikesRepository()
	cr := mock.GenerateNewMockedCategoriesRepository()
	commRep := mock.GenerateNewMockedCommentsRepository(ur)
	mocked := mock.GenerateNewMockedPostsRepository(cr, lr, ur)

	s := GenerateNewPostService(mocked, commRep)

	postId1 := uuid.New().String()

	post1 := entities.CreatePostServiceRequest{
		ID:          postId1,
		Title:       "o advento do código",
		Slug:        "code-advent",
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus. Quisque volutpat finibus nulla, ac semper augue lacinia quis. Sed vehicula neque a diam pulvinar, at tincidunt erat rutrum. Sed ultricies est at orci accumsan gravida. Aliquam quis tortor quis velit pretium commodo. Integer vel aliquet leo, eget egestas tortor. Quisque non tortor urna. Pellentesque dictum aliquet lectus, sagittis consequat ante pellentesque a. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus massa lorem, scelerisque vitae consectetur sed, luctus vel risus. ",
		Intro:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam ut eros ante. Nullam a urna in turpis porta gravida. Quisque a neque vel magna aliquam volutpat. Nunc elit dui, dapibus nec elit et, egestas tincidunt metus.",
		CategoryID:  "categoria 1",
		PublishedAt: time.Now().UTC().Format("2006-01-02 03:04:05"),
	}

	s.Create(&post1)

	err := s.Delete(post1.ID)
	if err != nil {
		t.Error(err)
	}

	if len(mocked.DB) != 0 {
		t.Errorf("expected db length to be 0, but got %d", len(mocked.DB))
	}
}

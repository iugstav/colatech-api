package server

import (
	"github.com/gin-gonic/gin"
	"github.com/iugstav/colatech-api/internal/middleware"
	"github.com/iugstav/colatech-api/pkg/category"
	"github.com/iugstav/colatech-api/pkg/comment"
	"github.com/iugstav/colatech-api/pkg/post"
	"github.com/iugstav/colatech-api/pkg/user"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(r *gin.RouterGroup, db *sqlx.DB) {
	postsRepository := post.GenerateNewPostsRepository(db)
	commentsRepository := comment.GenerateNewCommentsRepository(db)
	categoriesRepository := category.GenerateNewCategoriesRepository(db)
	usersRepository := user.GenerateNewUsersRepository(db)

	postService := post.GenerateNewPostService(postsRepository)
	commentService := comment.GenerateNewCommentService(commentsRepository)
	categoryService := category.GenerateNewCategoryService(*categoriesRepository)
	userService := user.GenerateNewUserService(usersRepository)

	postHandler := post.GeneratePostHandler(postService)
	commentHandler := comment.GenerateCommentHandler(commentService)
	categoryHandler := category.GenerateCategoryHandler(categoryService)
	userHandler := user.GenerateUserHandler(userService)

	// Posts routes
	postRG := r.Group("/posts")
	postRG.POST("/create",
		middleware.IsAuthenticatedMiddleware(),
		middleware.IsAuthorMiddleware(),
		postHandler.CreatePostController())
	postRG.GET("/get/all",
		middleware.IsAuthenticatedMiddleware(),
		postHandler.GetAllPostsController())
	postRG.GET("/get/all/minified", postHandler.GetAllMinifiedPostsController())
	postRG.GET("/get/:id", postHandler.GetPostByIdController())
	postRG.PATCH("/update/content",
		middleware.IsAuthenticatedMiddleware(),
		middleware.IsAuthorMiddleware(),
		postHandler.UpdatePostContentController())
	postRG.PATCH("/upload/image",
		middleware.IsAuthorMiddleware(),
		postHandler.UploadImageController())
	postRG.DELETE("/delete/:id",
		middleware.IsAuthenticatedMiddleware(),
		middleware.IsAuthorMiddleware(),
		postHandler.DeletePostController())

	// Commenta routes
	commentRG := r.Group("/comments")
	commentRG.POST("/create",
		middleware.IsAuthenticatedMiddleware(),
		commentHandler.CreateCommentController())
	commentRG.GET("/all/from/:pid", commentHandler.GetAllCommetsFromAPostController())
	commentRG.PATCH("/update/content",
		middleware.IsAuthenticatedMiddleware(),
		commentHandler.UpdateCommentContentController())
	commentRG.DELETE("delete/:id",
		middleware.IsAuthenticatedMiddleware(),
		middleware.IsAuthorMiddleware(),
		commentHandler.DeleteCommentController())

	// Categories routes
	categoryRG := r.Group("/categories")
	categoryRG.POST("/create",
		middleware.IsAuthenticatedMiddleware(),
		middleware.IsAuthorMiddleware(),
		categoryHandler.CreateCategoryController())
	categoryRG.GET("/all", categoryHandler.GetAllCategoriesController())
	categoryRG.PATCH("/update",
		middleware.IsAuthenticatedMiddleware(),
		middleware.IsAuthorMiddleware(),
		categoryHandler.UpdateCategoryNameController())
	categoryRG.DELETE("/delete/:id",
		middleware.IsAuthenticatedMiddleware(),
		middleware.IsAuthorMiddleware(),
		categoryHandler.DeleteCategoryController())

	// account routes
	userRG := r.Group("/users")
	userRG.GET("/get/:id",
		middleware.IsAuthenticatedMiddleware(),
		userHandler.GetUserByIdController())
	userRG.PATCH("/upload/image")
	userRG.POST("/create", userHandler.CreateUserController())
	userRG.POST("/login", userHandler.AuthenticateUserController())
}

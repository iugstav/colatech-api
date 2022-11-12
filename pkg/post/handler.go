package post

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/iugstav/colatech-api/internal/imgconv"
)

type PostHandler struct {
	PostService IPostService
}

func GeneratePostHandler(service IPostService) *PostHandler {
	return &PostHandler{
		PostService: service,
	}
}

func (h *PostHandler) CreatePostController() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TRANSFORMAR EM JSON PQ SEPAREI O UPLOAD DE IMAGEM EM UM ENDPOINT
		var request CreatePostFromEndpoint
		postCommonId := uuid.NewString()

		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid form data",
			})

			return
		}

		data := CreatePostServiceRequest{
			ID:          postCommonId,
			Title:       request.Title,
			Slug:        request.Slug,
			Intro:       request.Intro,
			Content:     request.Content,
			CategoryID:  request.CategoryID,
			PublishedAt: request.PublishedAt,
		}

		response, serviceErr := h.PostService.Create(&data)
		if serviceErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": serviceErr.Error(),
			})
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": response,
		})
	}
}

func (h *PostHandler) UploadImageController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request UploadPostCoverImageFromEndpoint

		if err := c.ShouldBindWith(&request, binding.Form); err != nil {
			log.Println("err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid form data",
			})

			return
		}

		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
			return
		}

		bytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})

			log.Fatal(err)
			return
		}

		webpImage, err := imgconv.ToWebp(bytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			log.Fatal(err)
			return
		}

		//	TODO: Add image extension verification
		newName := request.PostID + ".webp"
		dir, directoryErr := os.Getwd()
		if directoryErr != nil {
			log.Fatal((directoryErr))

			return
		}

		imagePath := fmt.Sprintf("%s/assets/posts/%s", dir, newName)
		webpFile, err := os.Create(imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})

			log.Fatal(err)
			return
		}
		webpFile.Write(webpImage)

		data := UploadPostCoverImageRequest{
			File:                 webpFile,
			ID:                   request.PostID,
			NameToUpload:         newName,
			CoverImageStaticPath: imagePath,
		}

		responseErr := h.PostService.UploadImage(&data)
		if responseErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": responseErr.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "image successfully uploaded",
		})
	}
}

func (h *PostHandler) GetAllPostsController() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := h.PostService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
	}
}

func (h *PostHandler) GetAllMinifiedPostsController() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := h.PostService.GetAllMinified()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
	}
}

func (h *PostHandler) GetPostByIdController() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("id")

		response, err := h.PostService.GetById(postId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
	}
}

func (h *PostHandler) UpdatePostContentController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request UpdatePostContentDTO

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := h.PostService.UpdateContent(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusNoContent, gin.H{
			"message": "Successfully updated content",
		})
	}
}

func (h *PostHandler) DeletePostController() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("id")

		if err := h.PostService.Delete(postId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusNoContent, gin.H{
			"message": "Successfully deleted post",
		})
	}
}

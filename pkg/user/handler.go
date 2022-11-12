package user

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

type UserHandler struct {
	UserService IUserService
}

func GenerateUserHandler(service IUserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) CreateUserController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreateUserFromEndpoint
		userCommonId := uuid.NewString()

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid form data",
			})

			return
		}

		user := CreateUserServiceRequest{
			ID:              userCommonId,
			UserName:        request.UserName,
			FirstName:       request.FirstName,
			LastName:        request.LastName,
			Email:           request.Email,
			Password:        request.Password,
			ImageStaticPath: "",
			CreatedAt:       request.CreatedAt,
		}

		serviceErr := h.UserService.Create(&user)
		if serviceErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": serviceErr.Error(),
			})
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User successfully created",
		})
	}
}

func (h *UserHandler) UploadProfileImageController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request UploadProfileImageFromEndpoint

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

		newName := request.UserID + ".webp"
		dir, directoryErr := os.Getwd()
		if directoryErr != nil {
			log.Fatal((directoryErr))

			return
		}

		imagePath := fmt.Sprintf("%s/assets/users/%s", dir, newName)
		webpFile, err := os.Create(imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})

			log.Fatal(err)
			return
		}
		webpFile.Write(webpImage)

		data := UploadProfileImageRequest{
			File:                   webpFile,
			ID:                     request.UserID,
			NameToUpload:           newName,
			ProfileImageStaticPath: imagePath,
		}

		responseErr := h.UserService.UploadIMage(&data)
		if responseErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": responseErr.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "image successfully uploaded",
		})
	}
}

func (h *UserHandler) AuthenticateUserController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request AuthenticateUserServiceRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})

			return
		}

		response, err := h.UserService.Authenticate(&request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}

func (h *UserHandler) GetUserByIdController() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")

		response, err := h.UserService.GetById(userId)
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

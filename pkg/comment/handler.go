package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iugstav/colatech-api/internal/entities"
)

type CommentHandler struct {
	CommentService ICommentService
}

func GenerateCommentHandler(service ICommentService) *CommentHandler {
	return &CommentHandler{CommentService: service}
}

func (h *CommentHandler) CreateCommentController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request entities.CreateCommentFromEndpoint

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		comment := entities.CreateCommentServiceRequest(request)

		response, serviceErr := h.CommentService.Create(&comment)
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

func (h *CommentHandler) GetAllCommetsFromAPostController() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("pid")

		response, err := h.CommentService.GetAllFromAPost(postId)
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

func (h *CommentHandler) UpdateCommentContentController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request entities.Comment

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if serviceErr := h.CommentService.UpdateContent(&request); serviceErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": serviceErr.Error(),
			})

			return
		}

		c.JSON(http.StatusNoContent, gin.H{
			"message": "Successfully updated content",
		})
	}
}

func (h *CommentHandler) DeleteCommentController() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId := c.Param("id")

		if err := h.CommentService.Delete(commentId); err != nil {
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

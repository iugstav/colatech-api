package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryService ICategoryService
}

func GenerateCategoryHandler(service ICategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: service}
}

func (h *CategoryHandler) CreateCategoryController() gin.HandlerFunc {
	return func(c *gin.Context) {
		type CreateCategoryRequest struct {
			Name string `json:"name" binding:"required,min=4,max=28"`
		}

		var data CreateCategoryRequest

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		response, serviceErr := h.CategoryService.Create(data.Name)
		if serviceErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": serviceErr.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": response,
		})
	}
}

func (h *CategoryHandler) GetAllCategoriesController() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := h.CategoryService.GetAll()
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

func (h *CategoryHandler) UpdateCategoryNameController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category Category

		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if serviceErr := h.CategoryService.UpdateName(&category); serviceErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": serviceErr.Error(),
			})

			return
		}

		c.JSON(http.StatusNoContent, gin.H{
			"message": "Successfully updated name",
		})
	}
}

func (h *CategoryHandler) DeleteCategoryController() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId := c.Param("id")

		if err := h.CategoryService.Delete(categoryId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusNoContent, gin.H{
			"message": "Successfully deleted category",
		})
	}
}

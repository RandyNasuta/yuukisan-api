package controllers

import (
	"net/http"
	"yuukisan/yuukisan/database"
	"yuukisan/yuukisan/helpers"
	"yuukisan/yuukisan/models"
	"yuukisan/yuukisan/structs"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve categories",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var categoryList []structs.Category

	for _, category := range categories {
		categoryList = append(categoryList, structs.Category{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
			CreatedBy:   category.CreatedBy,
			CreatedAt:   category.CreatedAt,
			UpdatedBy:   category.UpdatedBy,
			UpdatedAt:   category.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    categoryList,
	})
}

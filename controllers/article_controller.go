package controllers

import (
	"net/http"
	"yuukisan/yuukisan/database"
	"yuukisan/yuukisan/helpers"
	"yuukisan/yuukisan/models"
	"yuukisan/yuukisan/structs"

	"github.com/gin-gonic/gin"
)

func FindArticleById(c *gin.Context) {
	id := c.Param("id")

	var article models.Article

	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Article not found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Article found",
		Data: structs.Article {
			Id: article.Id,
			CategoryId: article.CategoryId,
			Title: article.Title,
			Slug: article.Slug,
			Content: article.Content,
			Cover: article.Cover,
			Writer: article.Writer,
			CreatedBy: article.CreatedBy,
			CreatedAt: article.CreatedAt,
			UpdatedBy: article.UpdatedBy,
			UpdatedAt: article.UpdatedAt,
		},
	})
}
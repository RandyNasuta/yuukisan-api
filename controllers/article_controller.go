package controllers

import (
	"net/http"
	"yuukisan/yuukisan/database"
	"yuukisan/yuukisan/helpers"
	"yuukisan/yuukisan/models"
	"yuukisan/yuukisan/structs"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var req = structs.ArticleCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	article := models.Article{
		CategoryId:  req.CategoryId,
		Title:       req.Title,
		Slug:        req.Slug,
		Content:     req.Content,
		Cover:       req.Cover,
		ImageBase64: req.ImageBase64,
		Writer:      req.Writer,				
		CreatedBy:   req.CreatedBy,
		CreatedAt:   req.CreatedAt,
		UpdatedBy:   req.UpdatedBy,
		UpdatedAt:   req.UpdatedAt,
	}

	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create article",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Article created successfully",
		Data: structs.Article{
			Id:          article.Id,
			CategoryId:  article.CategoryId,
			Title:       article.Title,
			Slug:        article.Slug,
			Content:     article.Content,
			Cover:       article.Cover,
			ImageBase64: article.ImageBase64,
			Writer:      article.Writer,
			CreatedBy:   article.CreatedBy,
			CreatedAt:   article.CreatedAt,
			UpdatedBy:   article.UpdatedBy,
			UpdatedAt:   article.UpdatedAt,
		},
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article

	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Article not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req = structs.ArticleUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	article.CategoryId = req.CategoryId
	article.Title = req.Title
	article.Slug = req.Slug
	article.Content = req.Content
	article.Cover = req.Cover
	article.ImageBase64 = req.ImageBase64
	article.Writer = req.Writer
	article.UpdatedBy = req.UpdatedBy
	article.UpdatedAt = req.UpdatedAt

	if err := database.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update article",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Article updated successfully",
		Data: structs.Article{
			Id:          article.Id,
			CategoryId:  article.CategoryId,
			Title:       article.Title,
			Slug:        article.Slug,
			Content:     article.Content,
			Cover:       article.Cover,
			ImageBase64: article.ImageBase64,
			Writer:      article.Writer,
			CreatedBy:   article.CreatedBy,
			CreatedAt:   article.CreatedAt,
			UpdatedBy:   article.UpdatedBy,
			UpdatedAt:   article.UpdatedAt,
		},
	})
}

func GetAllArticles(c *gin.Context) {
	var articles []models.Article

	if err := database.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve articles",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var articleList []structs.Article

	for _, article := range articles {
		articleList = append(articleList, structs.Article{
			Id:          article.Id,
			CategoryId:  article.CategoryId,
			Title:       article.Title,
			Slug:        article.Slug,
			Content:     article.Content,
			Cover:       article.Cover,
			ImageBase64: article.ImageBase64,
			Writer:      article.Writer,
			CreatedBy:   article.CreatedBy,
			CreatedAt:   article.CreatedAt,
			UpdatedBy:   article.UpdatedBy,
			UpdatedAt:   article.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Articles retrieved successfully",
		Data:    articleList,
	})
}

func FindArticleById(c *gin.Context) {
	id := c.Param("id")

	var article models.Article

	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Article not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Article found",
		Data: structs.Article{
			Id:          article.Id,
			CategoryId:  article.CategoryId,
			Title:       article.Title,
			Slug:        article.Slug,
			Content:     article.Content,
			Cover:       article.Cover,
			ImageBase64: article.ImageBase64,
			Writer:      article.Writer,
			CreatedBy:   article.CreatedBy,
			CreatedAt:   article.CreatedAt,
			UpdatedBy:   article.UpdatedBy,
			UpdatedAt:   article.UpdatedAt,
		},
	})
}

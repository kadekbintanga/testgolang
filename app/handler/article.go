package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"testgolang/app/models"
	"testgolang/app/repository"
	"testgolang/app/resource"
	"testgolang/app/helpers"
	"fmt"

)

type ArticleHandler struct{
	repo repository.ArticleRepository
}

func NewArticleHandler() *ArticleHandler{
	return &ArticleHandler{
		repository.NewArticleRepository(),
	}
}

func(h *ArticleHandler) HealthArticle(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Article Handler is ready!",
	})
}


func(h *ArticleHandler) CreateArticle(c *gin.Context){
	var req resource.InputArticle
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Print(err)
		errors := helpers.FormatValidationErrorInput(err)
		errorMessage := gin.H{"Error message": errors}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	Article := models.Article{
		Author: req.Author,
		Title: req.Title,
		Body: req.Body,
	}
	repo := h.repo
	res, err := repo.CreateArticle(Article)
	if err != nil {
		fmt.Print(err)
		errors := helpers.FormatValidationErrorSql(err)
		errorMessage := gin.H{"Error message": errors}
		response := helpers.APIResponse("Bad request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"id": res.ID,
		"author": res.Author,
		"title": res.Title,
		"body": res.Body,
	}

	response := helpers.APIResponse("Success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)



}

func (h *ArticleHandler) GetArticle(c *gin.Context){
	var author string = c.DefaultQuery("author", "")
	var search string = c.DefaultQuery("search", "")

	repo := h.repo

	res, err := repo.GetArticle(author, search)
	if err != nil {
		fmt.Print(err)
		errorMessage := gin.H{"Error message": "Something went wrong"}
		response := helpers.APIResponse("Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("Success", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}
package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
	// "testgolang/app/models"
	"testgolang/app/repository"
	// "testgolang/app/resource"

)

type ArticleHandler struct{
	repo repository.ArticleRepository
}

func NewArticleHandler() *ArticleHandler{
	return &ArticleHandler{
		repository.NewArticleRepository(),
	}
}

func HealthArticle(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Article Handler is ready!",
	})
}
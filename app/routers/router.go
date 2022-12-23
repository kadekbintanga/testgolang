package routers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"testgolang/app/handler"
)

func InitRouter(){
	ArticleHandler := handler.NewArticleHandler()
	r := gin.Default()
	api := r.Group("api/v1")

	api.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "I am ready!",
		})
	})

	api.GET("/article/health", ArticleHandler.HealthArticle)
	api.POST("/article", ArticleHandler.CreateArticle)
	api.GET("/article", ArticleHandler.GetArticle)


	r.Run()
}
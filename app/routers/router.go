package routers

import(
	"github.com/gin-gonic/gin"
	"github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"net/http"
	"testgolang/app/handler"
	"time"
)

func InitRouter(){
	ArticleHandler := handler.NewArticleHandler()
	r := gin.Default()
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	api := r.Group("api/v1")

	api.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "I am ready!",
		})
	})

	api.GET("/article/health", ArticleHandler.HealthArticle)
	api.POST("/article", ArticleHandler.CreateArticle)
	api.GET("/article", cache.CacheByRequestURI(memoryStore, 60*time.Second), ArticleHandler.GetArticle)


	r.Run()
}
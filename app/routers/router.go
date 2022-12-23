package routers

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(){
	r := gin.Default()
	api := r.Group("api/v1")

	api.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "I am ready!",
		})
	})


	r.Run()
}
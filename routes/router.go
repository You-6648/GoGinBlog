package routes

import (
	"GoGinBlog/until"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() {
	gin.SetMode(until.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		})
	}
	r.Run(until.HttpPort)

}

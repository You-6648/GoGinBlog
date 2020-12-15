package routes

import (
	"GoGinBlog/until"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() {
	gin.SetMode(until.AppMode)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		})
	}
	//s := &http.Server{
	//	Addr: fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler: r,
	//	ReadTimeout: setting.ReadTimeout,
	//	WriteTimeout: setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
	r.Run(until.HttpPort)

}

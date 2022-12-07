package routes

import (
	"github.com/gin-gonic/gin"
	"crud/controllers"
)

func UserRouter(router *gin.Engine) {
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello Word!s")
	// })

	router.POST("/",controllers.Adduser)
	router.GET("/find",controllers.FindUser)
	router.POST("/add:name/*action",controllers.Adduser)
	router.POST("/upload",controllers.Up)
	
}
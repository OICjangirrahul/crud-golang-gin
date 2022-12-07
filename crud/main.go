package main

import (

    "crud/config"
	"crud/routes"

	

    "github.com/gin-gonic/gin"
)


func main() {
	router := gin.New();
	config.Connect()
    router.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("index.html")
	routes.UserRouter(router)
	router.Run(":8000")
	


}
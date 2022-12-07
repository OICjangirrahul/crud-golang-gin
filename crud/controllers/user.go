package controllers

import (
	"crud/config"
	f "crud/functions"
	"crud/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Getuser(c *gin.Context) {
	// c.String(200, "hello rahul")
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"satus": 200, "users": users})
}

func Adduser(c *gin.Context) {
	var user models.User
	f.Errr(&user, c)
	config.DB.Create(&user)

}

func FindUser(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
	return

	//    // Get model if exist
	// 	var user models.User
	// 	name := c.Query("name")
	// 	if err := config.DB.Where("name = ?",name).First(&user).Error; err != nil {
	// 	 c.JSON(http.StatusBadRequest, gin.H{"status":400,"error": "データが見つかりませんでした"})
	// 	 return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"status":200,"data": &user})
}

func Up(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("file/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	filepath := c.Request.Host + c.Request.URL.Path + "/file" + "/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})

}

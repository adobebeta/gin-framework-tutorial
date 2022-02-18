package main

import (
	"gin-framework/controllers"
	"gin-framework/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.GET("book/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	r.Run()
}

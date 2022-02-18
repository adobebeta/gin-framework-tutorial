package controllers

import (
	"gin-framework/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET /books
//Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

//GET /books/:id
//Find a book by id
func FindBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

//POST /books
//Create new book
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	//Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Create Book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

//UPDATE /books/:id
//Update a book
func UpdateBook(c *gin.Context) {
	//Get model if exist
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	//validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&book).Update(input)
	c.JSON(http.StatusOK, gin.H{"date": book})
}

//DELETE /books/:id
//Delete a book
func DeleteBook(c *gin.Context) {
	//Get Model if exist
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	//delete
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": "delete successfully"})
}

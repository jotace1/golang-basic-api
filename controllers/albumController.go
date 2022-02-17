package controllers

import (
	"errors"
	"example/web-service-gin/db"
	"example/web-service-gin/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbuns(c *gin.Context) {
	var albums []entities.Album

	result := db.DB.Find(&albums)

	errors.Is(result.Error, gorm.ErrRecordNotFound)

	fmt.Println(albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func AddAlbum(c *gin.Context) {
	var newAlbum entities.Album

	c.BindJSON(&newAlbum)

	db.DB.Create(&newAlbum)

	c.IndentedJSON(http.StatusOK, newAlbum)

}

func FindById(c *gin.Context) {
	id := c.Param("id")

	var album entities.Album

	result := db.DB.First(&album, id)

	if err := errors.Is(result.Error, gorm.ErrRecordNotFound); err != false {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Album not found",
		})

		return
	}

	c.IndentedJSON(http.StatusOK, album)

}

func UpdateAlbum(c *gin.Context) {
	albumId := c.Param("id")
	var album entities.Album

	result := db.DB.First(&album, albumId)

	if err := errors.Is(result.Error, gorm.ErrRecordNotFound); err != false {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Album not found",
		})
		return
	}

	if err := c.BindJSON(&album); err != nil {
		fmt.Println(err.Error())
		return
	}

	db.DB.Save(&album)

	c.IndentedJSON(http.StatusOK, album)

}

func DeleteAlbum(c *gin.Context) {
	albumId := c.Param("id")

	var album entities.Album

	result := db.DB.First(&album, albumId)

	if err := errors.Is(result.Error, gorm.ErrRecordNotFound); err != false {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Album not found",
		})

		return
	}
	db.DB.Delete(&album)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Album successful deleted!",
	})
}

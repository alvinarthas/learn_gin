package routes

import (
	"github.com/alvinarthas/learn_gin/config"
	"github.com/alvinarthas/learn_gin/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

// GetHome buat
func GetHome(c *gin.Context) {

	items := []models.Article{}
	config.DB.Find(&items)

	// Return JSON
	c.JSON(200, gin.H{
		"status":  "berhasil",
		"message": "Berhasil Akses Home",
		"data":    items,
	})
}

// GetArticle buat
func GetArticle(c *gin.Context) {

	// Get Parameter
	slug := c.Param("slug")

	var item models.Article

	if config.DB.First(&item, "slug =  ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status":  "berhasil",
		"message": slug,
	})
}

// PostArticle buat
func PostArticle(c *gin.Context) {
	// Get Form
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil store data",
		"data":   item,
	})
}

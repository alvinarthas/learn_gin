package routes

import (
	"fmt"
	"strconv"
	"time"

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
		"status": "berhasil",
		"data":   item,
	})
}

// GetArticleByTag buat
func GetArticleByTag(c *gin.Context) {

	// Get Parameter
	tag := c.Param("tag")

	items := []models.Article{}

	if config.DB.Find(&items, "tag LIKE ?", "%"+tag+"%").RecordNotFound() || len(items) == 0 {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   items,
	})
}

// PostArticle buat
func PostArticle(c *gin.Context) {

	// Initialize Model
	oldItem := []models.Article{}
	// Get Parameter
	slug := slug.Make(c.PostForm("title"))

	// Do Query
	config.DB.First(&oldItem, "slug = ?", slug)

	if len(oldItem) >= 1 {
		slug = slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
		fmt.Println(slug)
	}

	// Get Form
	item := models.Article{
		Title:  c.PostForm("title"),
		Desc:   c.PostForm("desc"),
		Tag:    c.PostForm("tag"),
		Slug:   slug,
		UserID: uint(c.MustGet("jwt_user_id").(float64)),
	}

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil store data",
		"data":   item,
	})
}

// UpdateArticle buat update
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var item models.Article

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	if uint(c.MustGet("jwt_user_id").(float64)) != item.UserID {
		c.JSON(403, gin.H{
			"status":  "error",
			"message": "this data is forbidden"})
		c.Abort()
		return
	}

	config.DB.Model(&item).Where("id = ?", id).Updates(models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Tag:   c.PostForm("tag"),
	})

	c.JSON(200, gin.H{
		"status": "berhasil update data",
		"data":   item,
	})
}

// DeleteArticle is to delete Article
func DeleteArticle(c *gin.Context) {

	id := c.Param("id")
	var article models.Article

	config.DB.Where("id = ?", id).Delete(&article)
	c.JSON(200, gin.H{
		"status": "berhasil delete",
		"data":   article,
	})
}

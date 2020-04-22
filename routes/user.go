package routes

import (
	"github.com/alvinarthas/learn_gin/config"
	"github.com/alvinarthas/learn_gin/models"
	"github.com/gin-gonic/gin"
)

// GetProfile buat get seluruh artikel terkait user tersebut
func GetProfile(c *gin.Context) {

	userName := c.Param("user")

	var user models.User

	if config.DB.First(&user, "user_name =  ?", userName).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	item := config.DB.Where("id = ?", user.ID).Preload("Articles", "user_id = ?", user.ID).Find(&user)

	// Return JSON
	c.JSON(200, gin.H{
		"status":  "berhasil",
		"message": "Berhasil ke halaman Profile",
		"data":    item,
	})
}

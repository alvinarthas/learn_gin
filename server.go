package main

import (
	"github.com/alvinarthas/learn_gin/config"
	"github.com/alvinarthas/learn_gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	// set up database
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	//  Setting Default Router
	router := gin.Default()

	// Verison Grouping
	apiV1 := router.Group("/api/v1/")
	{
		apiV1.GET("/auth/:provider", routes.RedirectHandler)
		apiV1.GET("/auth/:provider/callback", routes.CallbackHandler)

		// Module Grouping
		articles := apiV1.Group("/article")
		{
			// Route initialization
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", routes.PostArticle)
		}
		// Old Route initialization
		// apiV1.GET("/", getHome)
		// apiV1.GET("/article/:title", getArticle)
		// apiV1.POST("/articles", postArticle)
	}

	router.Run()
}

package main

import (
	"github.com/NoNamePL/semaphore-demo-go-gin/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// load HTML templates
	router.LoadHTMLGlob("web/templates/*")
	// load css and js
	router.Static("/web/static/", "./web/static/")
	// Handle Index page
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/new/:article_id",handlers.GetArticle)
	router.Run(":8080")
}

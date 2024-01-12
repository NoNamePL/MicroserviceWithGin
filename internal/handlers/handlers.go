package handlers

import (
	"net/http"
	"strconv"

	"github.com/NoNamePL/semaphore-demo-go-gin/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(ctx *gin.Context) {

	// Get All articles
	articles := models.GetAllArticles()

	// Call the HTML method of the context to render a template
	ctx.HTML(
		// set the status ok (200)
		http.StatusOK,
		// Use the "index.html" template
		"index.html",
		//  pass the date that the page use
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func GetArticle(ctx *gin.Context) {
	articleId := ctx.Param("article_id") 
	article,err := models.GetArticleById(strconv.Atoi(articleId))

	ctx.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":article.Title,
			"payload":article,
		}
	)
	
}
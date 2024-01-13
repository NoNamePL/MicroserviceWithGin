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
	// Get article from url
	articleId := ctx.Param("article_id") 
	// Check if the article ID is valide
	if articleIdInt, err := strconv.Atoi(articleId); err == nil {
		// Check if the article exists
		if article,err := models.GetArticleById(articleIdInt); err == nil {
			ctx.HTML(
				// set http status ok (200)
				http.StatusOK,
				// Use the article.html template
				"article.html",
				// Pass the date that the page use
				gin.H{
					"title":article.Title,
					"payload":article,
				},
			)
		} else {
			// if the article is not found
			ctx.AbortWithError(http.StatusNotFound,err)
		}
	} else {
		// if article with this ID is not found
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}
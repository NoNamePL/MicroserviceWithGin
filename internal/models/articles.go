package models

import "errors"

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Articlelist = []Article{
	Article{Id: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{Id: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []Article {
	return Articlelist
}

func GetArticleById(id int) (*Article, error) {
	for _, val := range Articlelist {
		if val.Id == id {
			return &val, nil
		}
	}
	return nil, errors.New("Article not found")
}

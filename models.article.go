package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type article struct {
	ID int `json:"id"`
	Title string `json:title`
	Content string `json:content`
}

var articleList = []article{
	article{ID: 1,Title: "Article1",Content: "It's the first article"},
	article{ID: 2,Title: "Article2",Content: "It's the second article"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func render(c *gin.Context, data gin.H, templateName string){
	switch c.Request.Header.Get("Accept"){
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])

	case "application/xml":
		c.XML(http.StatusOK, data["payload"])

	default:
		c.HTML(http.StatusOK, templateName, data)
}}
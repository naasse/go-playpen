package controller

import (
    "net/http"
    "encoding/json"
    articleReps "go-playpen/articles-representations"
)

// TODO: These are mocked Articles.
// Next step will be to investigate persistence and ORM
var Articles = []articleReps.Article {
    articleReps.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
}

// Handle requests to get the list of articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Articles)
}


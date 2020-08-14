package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    articleReps "go-playpen/articles-representations"
    restReps "go-playpen/rest-representations"
)

var Articles []articleReps.Article

func root(w http.ResponseWriter, r *http.Request) {
    links := []restReps.Link {
        restReps.Link{Rel: "getArticles", Method: "get", Uri: "/articles"},
    }
    api := restReps.Api{Links: links}
    json.NewEncoder(w).Encode(api)
    fmt.Println("Endpoint Hit: API Root")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: Articles List")
    json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
    http.HandleFunc("/", root)
    http.HandleFunc("/articles", getArticles)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    Articles = []articleReps.Article {
        articleReps.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}


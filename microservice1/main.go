package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "go-playpen/microservice1/representations"
)

var Articles []representations.Article

func root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "{\"msg\": \"Hello, world!\"}")
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
    Articles = []representations.Article {
        representations.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}

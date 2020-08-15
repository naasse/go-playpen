package main

import (
    "log"
    "net/http"
    "go-playpen/articles-service/controller"
    "go-playpen/articles-service/constants"
)

// Define the endpoint handlers
func addHandlers() {
    http.HandleFunc(constants.RootUri + "/", controller.GetRoot)
    http.HandleFunc(constants.GetArticlesFullUri + "/", controller.GetArticles)
}

// Start listening and serving requests
func start() {
    log.Fatal(http.ListenAndServe(":10000", nil))
}

// Start the application
func main() {
    addHandlers()
    start()
}


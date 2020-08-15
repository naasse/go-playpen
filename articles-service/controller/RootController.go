package controller

import (
    "net/http"
    "encoding/json"
    "go-playpen/articles-service/constants"
    restReps "go-playpen/rest-representations"
    articlesReps "go-playpen/articles-representations"
)

// Handle requests to the root endpoint
func GetRoot(w http.ResponseWriter, r *http.Request) {
    links := []restReps.Link {
        restReps.Link{Rel: articlesReps.GetArticlesRel, Method: http.MethodGet, Uri: constants.GetArticlesFullUri + "/"},
    }
    api := restReps.Api{Links: links}
    json.NewEncoder(w).Encode(api)
}


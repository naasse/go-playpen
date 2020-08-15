package representations

var GetArticlesRel = "getArticles"

type Article struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

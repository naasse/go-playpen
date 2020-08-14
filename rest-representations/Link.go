package representations

type Link struct {
    Rel string `json:"rel"`
    Method string `json:"method"`
    Uri string `json:"uri"`
}

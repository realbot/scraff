package scraff

type Ad struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type AdExtractor interface {
	Extract() []Ad
}

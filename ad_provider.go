package scraff

type Ad struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"descr"`
}

type AdProvider interface {
	ID() string
	Ads() (ads []Ad, err error)
}

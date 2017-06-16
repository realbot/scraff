package scraff

type Ad struct {
	Url         string
	Title       string
	Description string
	Provider    string
}

type AdProvider interface {
	ID() string
	Ads() (ads []Ad, err error)
}

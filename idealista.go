package scraff

import (
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type IdealistaAdProvider struct {
	id        string
	retriever AdRetriever
}

func NewIdealistaAdProvider(id string, url string) *IdealistaAdProvider {
	return &IdealistaAdProvider{
		id: id,
		retriever: AdRetriever{
			Url: url,
		},
	}
}

func (ia IdealistaAdProvider) ID() string {
	return ia.id
}

func (ia IdealistaAdProvider) Ads() (ads []Ad, err error) {
	html, err := ia.retriever.retrieve()
	if err == nil {
		ads, err = ia.extract(html)
	}
	return
}

func (ia IdealistaAdProvider) extract(s string) (ads []Ad, err error) {
	root, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return
	}
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil {
			return scrape.Attr(n.Parent, "class") == "item-info-container"
		}
		return false
	}
	ads = []Ad{}
	matches := scrape.FindAll(root, matcher)
	for _, m := range matches {
		ads = append(ads, Ad{
			Title: scrape.Attr(m, "title"),
			Url:   "https://www.idealista.it" + scrape.Attr(m, "href"),
		})
	}
	return
}

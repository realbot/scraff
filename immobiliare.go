package scraff

import (
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ImmobiliareAdProvider struct {
	id        string
	retriever AdRetriever
}

func NewImmobiliareAdProvider(id string, url string) *ImmobiliareAdProvider {
	return &ImmobiliareAdProvider{
		id: id,
		retriever: AdRetriever{
			Url: url,
		},
	}
}

func (ia ImmobiliareAdProvider) ID() string {
	return ia.id
}

func (ia ImmobiliareAdProvider) Ads() (ads []Ad, err error) {
	html, err := ia.retriever.retrieve()
	if err == nil {
		ads, err = ia.extract(html)
	}
	return
}

func (ia ImmobiliareAdProvider) extract(s string) (ads []Ad, err error) {
	root, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return
	}
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "listing-item_body--content"
		}
		return false
	}
	ads = []Ad{}
	matches := scrape.FindAll(root, matcher)
	for _, m := range matches {
		ads = append(ads, Ad{
			Title: scrape.Text(m),
			Url:   scrape.Attr(m, "href"),
		})
	}
	return
}

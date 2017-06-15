package scraff

import (
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ImmobiliareAd struct {
	Retriever AdRetriever
}

func (ia ImmobiliareAd) ID() string {
	return "Immobiliare.it"
}

func (ia ImmobiliareAd) Ads() (ads []Ad, err error) {
	return nil, nil
}

func (ia ImmobiliareAd) extract(s string) (ads []Ad, err error) {
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

package scraff

import (
	"fmt"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ImmobiliareAd struct {
}

func (ia ImmobiliareAd) Extract(s string) (ads []Ad, err error) {
	root, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return
	}
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "athing"
		}
		return false
	}
	matches := scrape.FindAll(root, matcher)
	for i, m := range matches {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(m), scrape.Attr(m, "href"))
	}
	return ads, nil
}

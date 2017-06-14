package scraff

import (
	"io/ioutil"
	"net/http"
)

type AdRetriever struct {
	Url string
}

func (ar AdRetriever) retrieve() (html string, err error) {
	resp, err := http.Get(ar.Url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	html = string(data)
	return
}

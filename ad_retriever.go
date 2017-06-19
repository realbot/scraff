package scraff

import (
	"io/ioutil"
	"net/http"
)

type AdRetriever struct {
	Url string
}

func (ar AdRetriever) retrieve() (html string, err error) {
	/*on OSX??? verify...
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}*/
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

package scraff

import (
	"io/ioutil"
	"testing"
)

func TestExtractIdealista(t *testing.T) {
	ia := &IdealistaAdProvider{}
	content, err := ioutil.ReadFile("testdata/idealista.html")
	if err != nil {
		t.Error(err)
	}
	ads, err := ia.extract(string(content))
	if err != nil {
		t.Error(err)
	}
	/*for n, a := range ads {
		fmt.Printf("%d %s %s\n", n, a.Url, a.Title)
	}*/
	if len(ads) != 6 {
		t.Error("Wrong number of ads", ads)
	}
}

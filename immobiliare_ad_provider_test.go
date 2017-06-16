package scraff

import (
	"io/ioutil"
	"testing"
)

func TestExtract(t *testing.T) {
	ia := &ImmobiliareAdProvider{}
	content, err := ioutil.ReadFile("testdata/immobiliare.html")
	if err != nil {
		t.Error(err)
	}
	ads, err := ia.extract(string(content))
	if err != nil {
		t.Error(err)
	}
	if len(ads) != 5 {
		t.Error("Wrong number of ads", ads)
	}
}

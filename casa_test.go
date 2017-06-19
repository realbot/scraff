package scraff

import (
	"io/ioutil"
	"testing"
)

func TestExtractCasa(t *testing.T) {
	ia := &CasaAdProvider{}
	content, err := ioutil.ReadFile("testdata/casait.html")
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

package scraff

import (
	"github.com/golang/glog"
)

const processorVersion = "1.0.0"

type AdProcessor struct {
	Providers []AdProvider
	Store     AdStore
}

func (ap AdProcessor) Run() {
	glog.Infof("ad processor version %s", processorVersion)
	for {
		glog.Infof("Checking...")
		for _, provider := range ap.Providers {
			ads, err := provider.Ads()
			if err != nil {
				glog.Warningf("%s: %s", provider.ID, err)
			}
			ap.checkForNewAds(ads)
		}
	}
}

func (ap AdProcessor) checkForNewAds(ads []Ad) []Ad {
	return nil
}

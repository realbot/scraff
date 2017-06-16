package scraff

import (
	"time"

	"github.com/golang/glog"
)

const processorVersion = "1.0.0"

type AdProcessor struct {
	providers     []AdProvider
	store         AdStore
	sender        AdSender
	sleepDuration time.Duration
}

func NewAdProcessor(providers []AdProvider, store AdStore, sender AdSender, sleepDuration time.Duration) *AdProcessor {
	return &AdProcessor{
		providers:     providers,
		store:         store,
		sender:        sender,
		sleepDuration: sleepDuration,
	}
}

func (ap AdProcessor) Run() {
	glog.Infof("ad processor version %s", processorVersion)
	for {
		glog.Info("Checking...")
		var newads = []Ad{}
		for _, provider := range ap.providers {
			glog.Info("Checking " + provider.ID())
			ads, err := provider.Ads()
			if err != nil {
				glog.Warningf("%s: %s", provider.ID, err)
			} else {
				newads = append(newads, ap.checkForNewAds(ads)...)
			}
		}
		if len(newads) > 0 {
			ap.sender.Send(newads)
		}
		glog.Info("Done.")
		time.Sleep(ap.sleepDuration)
	}
}

func (ap AdProcessor) checkForNewAds(ads []Ad) (newads []Ad) {
	newads = []Ad{}
	for _, a := range ads {
		missing, err := ap.store.IsMissing(a)
		if err != nil {
			glog.Warning(err)
		} else if missing {
			newads = append(newads, a)
			err = ap.store.Add(a)
			if err != nil {
				glog.Warning(err)
			}
		}
	}
	return
}

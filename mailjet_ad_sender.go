package scraff

import (
	"github.com/golang/glog"
)

type MailAdSender struct {
}

func (m MailAdSender) Send(ads []Ad) (err error) {
	glog.Info("Sending...")
	for _, a := range ads {
		glog.Info("Sending %s\n", a.Url)
	}
	glog.Info("Done.")
	return
}

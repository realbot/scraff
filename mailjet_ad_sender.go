package scraff

import (
	"fmt"

	"github.com/golang/glog"
	. "github.com/mailjet/mailjet-apiv3-go"
)

type MailjetAdSender struct {
	mjPublicKey  string
	mjPrivateKey string
}

func NewMailjetAdSender(mjPublicKey string, mjPrivateKey string) *MailjetAdSender {
	return &MailjetAdSender{
		mjPublicKey:  mjPublicKey,
		mjPrivateKey: mjPrivateKey,
	}
}

func (m MailjetAdSender) Send(ads []Ad) (err error) {
	glog.Info("Sending...")

	subject := fmt.Sprintf("Found %d new ads", len(ads))
	body := "New ads found:<ul>"
	for _, a := range ads {
		body = body + fmt.Sprintf("<li><a href='%s'>%s</a>", a.Url, a.Url)
	}
	body = body + "</ul>"

	mailjetClient := NewMailjetClient(m.mjPublicKey, m.mjPrivateKey)
	email := &InfoSendMail{
		FromEmail: "mobile@realbot.me",
		FromName:  "Scraff Bot",
		Subject:   subject,
		HTMLPart:  body,
		Recipients: []Recipient{
			Recipient{
				Email: "mobile@realbot.me",
			},
		},
	}
	_, err = mailjetClient.SendMail(email)
	glog.Info("Done.")
	return
}

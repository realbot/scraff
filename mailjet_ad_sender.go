package scraff

import (
	"fmt"
)

type MailAdSender struct {
}

func (m MailAdSender) Send(ads []Ad) (err error) {
	fmt.Println("Sending...")
	for _, a := range ads {
		fmt.Printf("Sending %s\n", a.Url)
	}
	return
}

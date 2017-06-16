package scraff

type AdSender interface {
	Send(ads []Ad) (err error)
}

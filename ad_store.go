package scraff

type AdStore interface {
	IsMissing(ad Ad) (bool, error)
	Add(ad Ad) error
}

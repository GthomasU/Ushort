package storage

type StorageI interface {
	UpdateUrl(urlId, originalUrl string) bool
	SaveNewUrl(urlId, originalUrl string) bool
	GetOriginalUrl(urlId string) (string, error)
	RemoveOriginalUrl(urlId string) bool
}

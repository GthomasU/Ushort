package storage

type StorageI interface {
	UpdateUrl(urlId, originalUrl string) bool
	SaveNewUrl(urlId, originalUrl string) bool
	GetOriginalUrl(urlId string) (string, Error)
	RemoveOriginalUrl(urlId string) bool
}

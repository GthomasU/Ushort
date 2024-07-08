package tests

import (
	"Ushort/storage"
)

type StorageMock struct {
}

func (sm StorageMock) UpdateUrl(urlId, originalUrl string) bool {
	return true
}
func (sm StorageMock) SaveNewUrl(urlId, originalUrl string) bool {
	return true
}
func (sm StorageMock) GetOriginalUrl(urlId string) (string, storage.Error) {
	if urlId == "abcdefghi" {
		return "https://www.wikipedia.org", nil
	} else {
		return "", storage.RecordNotFound{}
	}
}
func (sm StorageMock) RemoveOriginalUrl(urlId string) bool {
	if urlId == "abcdefghi" {
		return true
	}
	return false
}

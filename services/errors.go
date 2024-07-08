package services

import "fmt"

type Error interface {
	Error() string
}

type UrlNotFound struct{}
type InvalidLength struct {
	Length int
}
type InvalidUrlId struct {
}

func (o UrlNotFound) Error() string {
	return "url not found"
}

func (o InvalidUrlId) Error() string {
	return "urlId is invalid"
}
func (e InvalidLength) Error() string {
	return fmt.Sprintf("Invalid Length: %v \n Length must be >0", e.Length)
}

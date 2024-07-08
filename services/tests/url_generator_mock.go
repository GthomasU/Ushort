package tests

import (
	"Ushort/services"
	"fmt"
)

type UrlGeneratorMock struct{}

func (ug UrlGeneratorMock) CreateRandomString(n int) (*string, services.Error) {
	mock := fmt.Sprint("abcdefghi")
	return &mock, nil
}

package services

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type UrlGenerator struct {
	Source rand.Source
}

func NewUrlGenerator() UrlGenerator {
	return UrlGenerator{
		Source: rand.NewSource(time.Now().UnixNano()),
	}
}

func (ug UrlGenerator) CreateRandomString(n int) (*string, Error) {
	if n <= 0 {
		return nil, InvalidLength{Length: n}
	}
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, ug.Source.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = ug.Source.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return (*string)(unsafe.Pointer(&b)), nil
}

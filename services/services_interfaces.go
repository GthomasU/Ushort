package services

type IUrlGenerator interface {
	CreateRandomString(n int) (*string, Error)
}

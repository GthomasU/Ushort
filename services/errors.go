package services 

type error interface {
	Error() string
}

type UrlNotFound struct {
}

func (o UrlNotFound) Error() string {
	return "url not found"
}

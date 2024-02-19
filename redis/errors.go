package redis

type error interface {
	Error() string
}

type RecordNotFound struct {
}

func (o RecordNotFound) Error() string {
	return "url not found"
}
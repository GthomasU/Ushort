package api 

type PostShortUrl struct {
	Url string `json:"url"`
}

type ResponseShortUrl struct {
	ShortedUrl string `json:"url"`
}

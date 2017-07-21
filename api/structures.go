package api

type Response struct {
	Message string `json:"message"`
	URL     string `json:"url"`
	ID      int    `json:"id"`
}

type jsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

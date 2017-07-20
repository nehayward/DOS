package api

type Attack struct {
	ID       int64  `json:"id"`
	URL      string `json:"url"`
	Requests int64  `json:"requests"`
	Period   int64  `json:"period"`
	Method   string `json:"method"`
	Info     Info   `json:"info"`
}

type Info struct {
	Requests      int64 `json:"requests"`
	Success       int64 `json:"success"`
	NetworkFailed int64 `json:"network_failed"`
	BadFailed     int64 `json:"bad_failed"`
}

type Response struct {
	Message string `json:"message"`
	URL     string `json:"url"`
	ID      int64  `json:"id"`
}

type Attacks []Attack

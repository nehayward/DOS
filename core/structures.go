package core

type Attacks []Attack

type Attack struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Requests int64  `json:"requests"`
	Method   string `json:"method"`
	Info     *Info  `json:"info"`
	Worker   Worker `json:"-"`
}

type Info struct {
	Requests      int64 `json:"requests"`
	Success       int64 `json:"success"`
	NetworkFailed int64 `json:"network_failed"`
	BadFailed     int64 `json:"bad_failed"`
}

type Worker struct {
	Channel chan bool
	ID      int64
}

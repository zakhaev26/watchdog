package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var clients = make(map[chan string]bool)
var addClient = make(chan chan string)
var removeClient = make(chan chan string)
var messages = make(chan string)

func handleEvents(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	notify := w.(http.CloseNotifier).CloseNotify()

	messageChan := make(chan string)
	addClient <- messageChan

	defer func() {
		removeClient <- messageChan
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		select {
		case <-notify:
			return
		case msg := <-messageChan:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		}
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/events", handleEvents)

	go func() {
		for {
			select {
			case c := <-addClient:
				clients[c] = true
			case c := <-removeClient:
				delete(clients, c)
			case msg := <-messages:
				for c := range clients {
					c <- msg
				}
			}
		}
	}()

	go func() {
		for i := 0; ; i++ {
			messages <- fmt.Sprintf("message %d", i)
			time.Sleep(2 * time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", r))
}

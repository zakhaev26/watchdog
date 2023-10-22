package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)
var i int = 0;

func streamer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// the current implementation uses JSON encoding to send integers, which is not compatible with the SSE format. SSE expects data to be formatted with specific event tags
	// In SSE, the data is sent in a specific format, usually preceded by data:, and each event is terminated by a pair of newline characters (\n\n). This format is essential for the client to interpret the streamed data correctly.
	for ; ; i++ {
		by , _ := fmt.Fprintf(w, "data: %d\n\n", i)
		fmt.Println(by); //data: = 6 bytes , integer = 4bytes
		w.(http.Flusher).Flush()
		time.Sleep(3 * time.Second)
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/stream", streamer)

	log.Fatal(http.ListenAndServe(":8080", r))
}

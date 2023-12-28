package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/zakhaev26/hypernotifs/db"
	"github.com/zakhaev26/hypernotifs/metrics"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ScoreEvents struct {
	Scorer     string `json:"scorer,omitempty" bson:"scorer,omitempty"`
	Rival      string `json:"rival,omitempty" bson:"rival,omitempty"`
	ScorerTeam string `json:"scorer_team,omitempty" bson:"scorer_team,omitempty"`
	Time       string `json:"time,omitempty" bson:"time,omitempty"`
}

func main() {
	db.Init()
	r := mux.NewRouter()
	go metrics.LogCPUUsage()
	r.HandleFunc("/", SSEHandler)
	r.HandleFunc("/post/data", PostHandler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("/post")
	var temp ScoreEvents
	json.NewDecoder(r.Body).Decode(&temp)
	var wg sync.WaitGroup
	wg.Add(1)
	var inserted *mongo.InsertOneResult

	go func() {
		defer wg.Done()
		inserted = postData(temp)
		fmt.Println(inserted)
	}()

	json.NewEncoder(w).Encode("inserted")
	wg.Wait()

}

func postData(data ScoreEvents) *mongo.InsertOneResult {

	inserted, err := db.Collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted)
	return inserted
}

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Log In")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	changeStreamOptions := options.ChangeStream().
		SetFullDocument(options.UpdateLookup)

	changeStream, err := db.Collection.Watch(context.Background(), mongo.Pipeline{}, changeStreamOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer changeStream.Close(context.Background())

	messageChannel := make(chan bson.M)
	defer close(messageChannel)

	go func() {
		for changeStream.Next(context.Background()) {
			var changeDocument bson.M
			if err := changeStream.Decode(&changeDocument); err != nil {
				log.Fatal(err)
			}
			messageChannel <- changeDocument
			fmt.Println("Change document:", changeDocument)
		}
	}()

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		fmt.Println("Client disconnected")
	}()

	for {
		select {
		case msg := <-messageChannel:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-notify:
			fmt.Println("Client disconnected")
			return
		}
	}
}

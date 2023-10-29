package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/darthsalad/socketeer"
)

func main() {
	
	dbName := "test";
	collName  := "forms" 
	s, err := socketeer.NewSocketeer("mongodb+srv:soubhik:iN7fj94FdEXPn2Za@itb.b6ocgev.mongodb.net/?retryWrites=true&w=majority", dbName, collName)
	if err != nil {
		log.Fatal(err)
	}

	fields := []string{"title", "text"}
	url := "localhost:8080"
	endpoint := "/listen"

	s.Start(fields, url, endpoint)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sigCh

	s.Stop()
	fmt.Println("Socketeer stopped gracefully.")

	defer s.Stop()
}

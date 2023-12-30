package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/zakhaev26/elastic/consumer"
)

func main() {

	consumer, worker := consumer.Consumer("watchdog-critical-logs")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("%s\n", (msg.Value))
				message := `{ "index" : { "_index" : "nuakhai", "_id" : "` + strconv.Itoa(msgCount) + `" } }` +
					`
{"` + string(msg.Value) + `":` + `null}` +
					`

`
				err := os.WriteFile("reqs", []byte(message), 0755)
				if err != nil {
					fmt.Println("Error creating reqs:", err)
				}

				bashScript := []byte(
					`curl -XPOST -i -k \
					-H "Content-Type: application/x-ndjson" \
					-H "Authorization: ApiKey YnRQMnVZd0JiVGNVQXhUR2hsVGo6QjJRNndEQ2dSTHVmVjdtNGx1OV9GQQ==" \
					https://watchdog.es.asia-south1.gcp.elastic-cloud.com/_bulk --data-binary "@reqs"; echo      
					`)

				// Save the Bash script to a file
				err = os.WriteFile("myscript.sh", bashScript, 0755)
				if err != nil {
					fmt.Println("Error creating Bash script:", err)
					return
				}

				// Run the Bash script using the "sh" command
				cmd := exec.Command("sh", "myscript.sh")

				// Redirect standard output and error to the console
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				// Run the command
				err = cmd.Run()
				if err != nil {
					fmt.Println("Error running Bash script:", err)
					return
				}
			case <-sigchan:
				fmt.Println("Interrupt is detected")

				err := os.Remove("reqs")
				if err != nil {
					fmt.Println("Error removing LogGen:", err)
				}

				err = os.Remove("myscript.sh")
				if err != nil {
					fmt.Println("Error removing Bash script:", err)
				}
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}

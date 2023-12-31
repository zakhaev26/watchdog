package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/zakhaev26/elastic/consumer"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env")
		return
	}

	CRITICAL_LOG_NODE_ID := os.Getenv("CRITICAL_LOG_NODE_ID")
	ES_HOST := os.Getenv("ES_HOST")
	ES_API_KEY := os.Getenv("ES_API_KEY")

	consumer, worker := consumer.Consumer(CRITICAL_LOG_NODE_ID)

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
				var stat string = string(msg.Value)
				parts := strings.Split(stat, " ")

				cpuUsage, _ := strconv.ParseFloat(parts[0], 64)
				timeValue := parts[1]
				message := `{ "index" : { "_index" : "` + CRITICAL_LOG_NODE_ID + `", "_id" : "` + strconv.Itoa(msgCount) + `" } }
{"cpu_usage": ` + strconv.FormatFloat(cpuUsage, 'f', -1, 64) + `, "time": "` + timeValue + `"}` + "\n"

				fmt.Println("UH:", message)
				err := os.WriteFile("reqs", []byte(message), 0755)
				if err != nil {
					fmt.Println("Error creating reqs:", err)
				}

				bashScript := []byte(
					`curl -XPOST -i -k \
					-H "Content-Type: application/x-ndjson" \
					-H "Authorization: ApiKey ` + ES_API_KEY + `" \` +
						ES_HOST + `/_bulk	 --data-binary "@reqs"; echo      
					`)

				// Save the Bash script to a file
				err = os.WriteFile("critical.sh", bashScript, 0755)
				if err != nil {
					fmt.Println("Error creating Bash script:", err)
					return
				}

				// Run the Bash script using the "sh" command
				cmd := exec.Command("sh", "critical.sh")

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

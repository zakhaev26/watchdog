package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	kafkaProducer "github.com/zakhaev26/critical_producer/kafka"
	processor "github.com/zakhaev26/frequent_producer/cpu"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env")
		return
	}
	FREQUENT_LOG_NODE_ID := os.Getenv("FREQUENT_LOG_NODE_ID")

	for {
		data := processor.CpuUsageEachSecond()
		err := kafkaProducer.PushToKafka(FREQUENT_LOG_NODE_ID, data)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Produced.")
		}
		time.Sleep(time.Second * 2)
	}
}

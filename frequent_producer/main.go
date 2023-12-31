package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	kafkaProducer "github.com/zakhaev26/critical_producer/kafka"
	"github.com/zakhaev26/critical_producer/protobuf"
	processor "github.com/zakhaev26/frequent_producer/cpu"
	"google.golang.org/protobuf/proto"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env")
		return
	}
	FREQUENT_LOG_NODE_ID := os.Getenv("FREQUENT_LOG_NODE_ID")

	for {
		cpu, time_ := processor.CpuUsageEachSecond()
		// fString := data + " " + time_
		msg := &protobuf.KibanaMessage{
			CpuUsage:  cpu,
			Time:      time_,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		data, err := proto.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
		
		err = kafkaProducer.PushToKafka(FREQUENT_LOG_NODE_ID,string(data));

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Produced.")
		}
		time.Sleep(time.Second * 2)
	}
}

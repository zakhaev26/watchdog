package main

import (
	"fmt"
	"time"

	kafkaProducer "github.com/zakhaev26/critical_producer/kafka"
	processor "github.com/zakhaev26/frequent_producer/cpu"
)

func main() {
	for {
		data := processor.CpuUsageEachSecond()
		err := kafkaProducer.PushToKafka("watchdog-frequent-logs", data)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Produced.")
		}
		time.Sleep(time.Second*2);
	}
}

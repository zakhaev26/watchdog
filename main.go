package main

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zakhaev26/WatchDog/processor"
	"github.com/zakhaev26/WatchDog/producer"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			AvgCpuUsage := processor.FetchCpuUsage()

			if AvgCpuUsage > 80.0 {
				logEntry := logrus.WithFields(logrus.Fields{
					"cpu_usage": AvgCpuUsage,
				})

				logEntry.Info("Critical event detected")

				// Convert log entry fields to JSON
				jsonFields, err := json.Marshal(logEntry.Data)
				if err != nil {
					logEntry.WithError(err).Error("Failed to serialize log entry fields")
					continue
				}

				err = producer.PushToKafka("WDCriticalLogs", jsonFields)

				if err != nil {
					logEntry.WithError(err).Error("Failed to produce to Kafka")
				} else {
					logEntry.Info("Produced in WDCriticalLogs!")
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {

			data := processor.CpuUsageEachSecond()
			err := producer.PushToKafka("WDDailyLogs", []byte(data))

			if err != nil {
				logrus.WithError(err).Error("Failed to produce to Kafka")
			} else {
				logrus.Info("Produced in WDDailyLogs!")
			}

			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}

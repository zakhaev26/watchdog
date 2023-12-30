package criticalLogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zakhaev26/WatchDog/mailing"
	"github.com/zakhaev26/WatchDog/processor"
	"github.com/zakhaev26/WatchDog/producer"
)

func CriticalLog() {
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

			file, err := os.Open("files/index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			content, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return
			}

			// fmt.Println(string(content))

			val := mailing.SendMail(string(content), "Cpu OverClocking - Watchdog Metrics", "b422056@iiit-bh.ac.in")
			time.Sleep(time.Second * 60)
			fmt.Println(val)
		} else {
			continue
		}
	}
}

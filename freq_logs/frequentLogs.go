package frequentLogs

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zakhaev26/WatchDog/processor"
	"github.com/zakhaev26/WatchDog/producer"
)

func FrequentLogger() {
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
}

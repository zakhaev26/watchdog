package main

import (
	"sync"

	"github.com/sirupsen/logrus"
	criticalLogs "github.com/zakhaev26/WatchDog/critical_logs"
	frequentLogs "github.com/zakhaev26/WatchDog/freq_logs"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		criticalLogs.CriticalLog()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		frequentLogs.FrequentLogger()
	}()

	wg.Wait()
}

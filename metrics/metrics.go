package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func GetCPUUsage() float64 {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)

	// Calculate CPU usage as a percentage
	cpuUsage := 100 * float64(stat.Sys-stat.HeapReleased) / float64(stat.Sys)

	return cpuUsage
}


func LogCPUUsage() {
	// Open the log file for writing
	logFile, err := os.Create("cpu_usage.log")
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()
	// Write header to the log file
	logFile.WriteString("Time,CPU Usage (%)\n")

	// Log CPU usage every 5 seconds
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Get CPU usage
			cpuUsage := GetCPUUsage()

			// Get the current time
			currentTime := time.Now()

			// Format and write the log entry
			logEntry := fmt.Sprintf("[TIME]%s, [CPU]%f\n", currentTime.Format("2006-01-02 15:04:05"), cpuUsage)
			logFile.WriteString(logEntry)

			// Print to console (optional)
			fmt.Print(logEntry)
		}
	}
}

func main() {
	LogCPUUsage()
}


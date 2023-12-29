package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/cpu"
)

// Global variable to store CPU information
var cpuDetails []cpu.InfoStat

func main() {
	// Get CPU information
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}

	// Log and assign CPU information to the global variable
	fmt.Println("CPU Info:")
	for _, info := range cpuInfo {
		fmt.Printf("  CPU %v: %v\n", info.CPU, info.Model)

		// Append CPU details to the global variable
		cpuDetails = append(cpuDetails, info)
	}

	// // Monitor CPU usage for a few seconds
	// fmt.Println("\nMonitoring CPU Usage (Press Ctrl+C to stop):")
	// for {
	// 	// Get CPU usage percentages
	// 	cpuPercentages, err := cpu.Percent(time.Second, false)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Extract usage percentage into a variable
	// 	cpuUsage := cpuPercentages[0]
	// 	for i, data := range cpuPercentages {
	// 		fmt.Printf("x i:%d data : %v", i, data)
	// 	}
	// 	// Print CPU usage percentage
	// 	fmt.Printf("  Usage Percentage (CPU x): %.2f%%\n", cpuUsage)
	// 	time.Sleep(time.Second)
	// }

	for key, val := range cpuDetails {
		fmt.Printf("%v | %v\n", key, val)
	}

}

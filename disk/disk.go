package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/disk"
)

// Global variable to store disk information
var diskDetails []disk.PartitionStat

func main() {
	// Get disk partitions information
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Fatal(err)
	}

	// Log and assign disk partition information to the global variable
	fmt.Println("Disk Partitions:")
	for _, partition := range partitions {
		fmt.Printf("  Mountpoint: %s\n", partition.Mountpoint)
		fmt.Printf("    Device: %s\n", partition.Device)
		fmt.Printf("    Fstype: %s\n", partition.Fstype)

		// Append disk partition details to the global variable
		diskDetails = append(diskDetails, partition)
	}

	// Monitor disk usage for a few seconds
	fmt.Println("\nMonitoring Disk Usage (Press Ctrl+C to stop):")
	for i := 0; i < 5; i++ {
		// Get disk usage percentages
		diskUsage, err := disk.Usage(partitions[0].Mountpoint)
		if err != nil {
			log.Fatal(err)
		}

		// Print disk usage details
		fmt.Printf("  Disk Usage (Partition %d - %s): %.2f%%\n", i, partitions[0].Mountpoint, diskUsage.UsedPercent)
	}

	// Print global variable containing disk details
	fmt.Println("\nStored Disk Details:")

	for _, detail := range diskDetails {
		fmt.Println("--------------------")
		fmt.Printf("  Mountpoint: %s\n", detail.Mountpoint)
		fmt.Printf("    Device: %s\n", detail.Device)
		fmt.Printf("    Fstype: %s\n", detail.Fstype)
	
	}
}

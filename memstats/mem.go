package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	// Get memory information
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate percentages
	totalPercentage := 100.0
	freePercentage := (float64(memInfo.Free) / float64(memInfo.Total)) * totalPercentage
	usedPercentage := (float64(memInfo.Used) / float64(memInfo.Total)) * totalPercentage
	buffersPercentage := (float64(memInfo.Buffers) / float64(memInfo.Total)) * totalPercentage
	cachedPercentage := (float64(memInfo.Cached) / float64(memInfo.Total)) * totalPercentage

	// Print the results
	fmt.Println("Memory Info (Percentages):")
	fmt.Printf("  Total: %.2f%%\n", totalPercentage)
	fmt.Printf("  Free: %.2f%%\n", freePercentage)
	fmt.Printf("  Used: %.2f%%\n", usedPercentage)
	fmt.Printf("  Buffers: %.2f%%\n", buffersPercentage)
	fmt.Printf("  Cached: %.2f%%\n", cachedPercentage)

	// Additional fields
	fmt.Println("\nAdditional Fields:")
	fmt.Printf("  Active: %v MB\n", memInfo.Active)
	fmt.Printf("  Inactive: %v MB\n", memInfo.Inactive)
	fmt.Printf("  Wired: %v MB\n", memInfo.Wired)
	fmt.Printf("  Laundry: %v MB\n", memInfo.Laundry)
	fmt.Printf("  Writeback: %v MB\n", memInfo.Writeback)
	fmt.Printf("  Dirty: %v MB\n", memInfo.Dirty)
	fmt.Printf("  WritebackTmp: %v MB\n", memInfo.WritebackTmp)
	fmt.Printf("  Shared: %v MB\n", memInfo.Shared)
	fmt.Printf("  Slab: %v MB\n", memInfo.Slab)
	fmt.Printf("  SReclaimable: %v MB\n", memInfo.SReclaimable)
	fmt.Printf("  SUnreclaim: %v MB\n", memInfo.SUnreclaim)
	fmt.Printf("  PageTables: %v MB\n", memInfo.PageTables)
	fmt.Printf("  SwapCached: %v MB\n", memInfo.SwapCached)
	fmt.Printf("  CommitLimit: %v MB\n", memInfo.CommitLimit)
	fmt.Printf("  CommittedAS: %v MB\n", memInfo.CommittedAS)
	fmt.Printf("  HighTotal: %v MB\n", memInfo.HighTotal)
	fmt.Printf("  HighFree: %v MB\n", memInfo.HighFree)
	fmt.Printf("  LowTotal: %v MB\n", memInfo.LowTotal)
	fmt.Printf("  LowFree: %v MB\n", memInfo.LowFree)
	fmt.Printf("  SwapTotal: %v MB\n", memInfo.SwapTotal)
	fmt.Printf("  SwapFree: %v MB\n", memInfo.SwapFree)
	fmt.Printf("  Mapped: %v MB\n", memInfo.Mapped)
	fmt.Printf("  VMallocTotal: %v MB\n", memInfo.VMallocTotal)
	fmt.Printf("  VMallocUsed: %v MB\n", memInfo.VMallocUsed)
	fmt.Printf("  VMallocChunk: %v MB\n", memInfo.VMallocChunk)
	fmt.Printf("  HugePagesTotal: %v MB\n", memInfo.HugePagesTotal)
	fmt.Printf("  HugePagesFree: %v MB\n", memInfo.HugePagesFree)
	fmt.Printf("  HugePageSize: %v MB\n", memInfo.HugePageSize)
}

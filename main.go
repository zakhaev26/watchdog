package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/host"
)

func main() {
	infoStat, err := host.Info()

	if err != nil {
		log.Fatal(err)
	}

	var platformFamily string = infoStat.PlatformFamily
	fmt.Println(platformFamily)

	var hostID string = infoStat.HostID
	fmt.Println(hostID)

	var bootTime uint64 = infoStat.BootTime
	fmt.Println(bootTime)

	var kernelArch string = infoStat.KernelArch
	fmt.Println(kernelArch)

	var procs uint64 = infoStat.Procs
	fmt.Println(procs)

	var upTime string = formatUptime(infoStat.Uptime)

	fmt.Println(upTime)

	var vRole string = infoStat.VirtualizationRole
	fmt.Println(vRole)

	var pmVersion string = infoStat.PlatformVersion
	fmt.Println(pmVersion)

	var vSys string = infoStat.VirtualizationSystem
	fmt.Println(vSys)

}

func formatUptime(uptime uint64) string {
	duration := time.Duration(uptime) * time.Second
	return duration.String()
}

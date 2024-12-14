package monitor

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"		
	"time"
	"bytes"
	"os/exec"
	"fmt"
)

func GetCPUUsage() (float64, error) {		
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

func GetMemoryUsage() (float64, uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return vmStat.UsedPercent, vmStat.Used, nil
}

func GetDiskUsage() (float64, uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return diskStat.UsedPercent, diskStat.Used, nil
}

func GetNetworkStats() ([]net.IOCountersStat, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func GetCPUCount() (int, error) {
	return cpu.Counts(true)
}

func GetMemoryTotal() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Total, nil
}

func GetDiskTotal() (uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.Total, nil
}

func GetNetworkInterfaces() ([]net.InterfaceStat, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return interfaces, nil
}

func GetDiskReadWriteStats() (map[string]disk.IOCountersStat, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func GetCPUFrequency() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

func GetDiskInodes() (uint64, uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return inodeStat.InodesUsed, inodeStat.InodesFree, nil
}

func GetMemorySwapTotal() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Total, nil
}

func GetMemorySwapUsed() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Used, nil
}

func GetDiskReadBytes() (uint64, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return 0, err
	}
	var totalRead uint64
	for _, stat := range stats {
		totalRead += stat.ReadBytes
	}
	return totalRead, nil
}

func GetDiskWriteBytes() (uint64, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return 0, err
	}
	var totalWrite uint64
	for _, stat := range stats {
		totalWrite += stat.WriteBytes
	}
	return totalWrite, nil
}

func GetNetworkSentBytes() (uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}
	var totalSent uint64
	for _, stat := range stats {
		totalSent += stat.BytesSent
	}
	return totalSent, nil
}

func GetNetworkReceivedBytes() (uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}
	var totalReceived uint64
	for _, stat := range stats {
		totalReceived += stat.BytesRecv
	}
	return totalReceived, nil
}

func GetCPUPercentages(interval float64) ([]float64, error) {
	return cpu.Percent(time.Duration(interval * float64(time.Second)), true)
}

func GetDiskFree() (uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.Free, nil
}

func GetMemoryFree() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Free, nil
}

func GetCPUIdle() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return 100 - percentages[0], nil
}

func GetDiskInodesUsed() (uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return inodeStat.InodesUsed, nil
}

func GetDiskInodesFree() (uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return inodeStat.InodesFree, nil
}

func GetNetworkPacketsSent() (uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}
	var totalPacketsSent uint64
	for _, stat := range stats {
		totalPacketsSent += stat.PacketsSent
	}
	return totalPacketsSent, nil
}

func GetNetworkPacketsReceived() (uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}
	var totalPacketsReceived uint64
	for _, stat := range stats {
		totalPacketsReceived += stat.PacketsRecv
	}
	return totalPacketsReceived, nil
}

func GetSwapTotal() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Total, nil
}

func GetSwapUsed() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Used, nil
}

func GetSwapFree() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Free, nil
}

func GetCPUTemperature() (float64, error) {
	temperature, err := readCPUTemperature() 
	if err != nil {
		return 0, err
	}
	return temperature, nil
}

func readCPUTemperature() (float64, error) {
	cmd := exec.Command("sensors")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	output := out.String()
	var temperature float64
	_, err = fmt.Sscanf(output, "Core 0: +%f°C", &temperature)
	if err != nil {
		return 0, err
	}

	return temperature, nil
}

func GetDiskIOStats() (map[string]disk.IOCountersStat, error) {
	return disk.IOCounters()
}

func GetMemoryCached() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Cached, nil
}

func GetMemoryBuffers() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Buffers, nil
}
func GetDiskWriteBytesPerSecond() (uint64, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return 0, err
	}
	var totalWrite uint64
	for _, stat := range stats {
		totalWrite += stat.WriteBytes
	}
	return totalWrite / uint64(time.Now().Unix()), nil
}

func GetDiskReadBytesPerSecond() (uint64, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return 0, err
	}
	var totalRead uint64
	for _, stat := range stats {
		totalRead += stat.ReadBytes
	}
	return totalRead / uint64(time.Now().Unix()), nil 
}

func GetNetworkErrorStats() ([]net.IOCountersStat, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	var errorStats []net.IOCountersStat
	for _, stat := range stats {
		if stat.Errin > 0 || stat.Errout > 0 {
			errorStats = append(errorStats, stat)
		}
	}
	return errorStats, nil
}

func GetCPUSystemTime() (float64, error) {
	cpuStat, err := cpu.Times(false)
	if err != nil {
		return 0, err
	}
	if len(cpuStat) == 0 {
		return 0, nil
	}
	return cpuStat[0].System, nil
}

func GetCPUUserTime() (float64, error) {
	cpuStat, err := cpu.Times(false)
	if err != nil {
		return 0, err
	}
	if len(cpuStat) == 0 {
		return 0, nil
	}
	return cpuStat[0].User, nil
}

func GetDiskUsageByPath(path string) (float64, uint64, error) {
	diskStat, err := disk.Usage(path)
	if err != nil {
		return 0, 0, err
	}
	return diskStat.UsedPercent, diskStat.Used, nil
}

func GetMemorySwapFree() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Free, nil
}

func GetNetworkInterfaceStats(name string) (*net.InterfaceStat, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range interfaces {
		if iface.Name == name {
			return &iface, nil
		}
	}
	return nil, nil
}

func GetDiskInodeUsage() (uint64, uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return inodeStat.InodesUsed, inodeStat.InodesFree, nil
}

func GetMemoryTotalSwap() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Total, nil
}
func GetNetworkPacketsDropped() (uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}
	var totalDropped uint64
	for _, stat := range stats {
		totalDropped += stat.Dropin
		totalDropped += stat.Dropout
	}
	return totalDropped, nil
}

func GetDiskTotalInodes() (uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return inodeStat.InodesTotal, nil
}

func GetMemoryTotalUsed() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Used, nil
}

func GetDiskTotalUsed() (uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.Used, nil
}

package monitor

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
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

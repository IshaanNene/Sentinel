package monitor

import (
	"github.com/shirou/gopsutil/v3/net"		
)

func GetNetworkStats() ([]net.IOCountersStat, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func GetNetworkInterfaces() ([]net.InterfaceStat, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return interfaces, nil
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

func GetNetworkTotalErrors() (uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}
	var totalErrors uint64
	for _, stat := range stats {
		totalErrors += stat.Errin + stat.Errout
	}
	return totalErrors, nil
}
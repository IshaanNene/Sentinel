package monitor

import (
	"github.com/shirou/gopsutil/v3/disk"
	"time"
)

func GetDiskUsage() (float64, uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return diskStat.UsedPercent, diskStat.Used, nil
}

func GetDiskTotal() (uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.Total, nil
}

func GetDiskReadWriteStats() (map[string]disk.IOCountersStat, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func GetDiskInodes() (uint64, uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return inodeStat.InodesUsed, inodeStat.InodesFree, nil
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

func GetDiskFree() (uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.Free, nil
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

func GetDiskIOStats() (map[string]disk.IOCountersStat, error) {
	return disk.IOCounters()
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

func GetDiskUsageByPath(path string) (float64, uint64, error) {
	diskStat, err := disk.Usage(path)
	if err != nil {
		return 0, 0, err
	}
	return diskStat.UsedPercent, diskStat.Used, nil
}

func GetDiskInodeUsage() (uint64, uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return inodeStat.InodesUsed, inodeStat.InodesFree, nil
}

func GetDiskTotalInodes() (uint64, error) {
	inodeStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return inodeStat.InodesTotal, nil
}

func GetDiskTotalUsed() (uint64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.Used, nil
}

func GetDiskSerialNumber() (string, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return "", err
	}
	for _, stat := range stats {
		if stat.SerialNumber != "" {
			return stat.SerialNumber, nil
		}
	}
	return "", nil
}
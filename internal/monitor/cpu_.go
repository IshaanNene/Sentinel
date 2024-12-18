package monitor

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

func GetCPUUsage() (float64, error) {		
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

func GetCPUCount() (int, error) {
	return cpu.Counts(true)
}

func GetCPUFrequency() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

func GetCPUPercentages(interval float64) ([]float64, error) {
	return cpu.Percent(time.Duration(interval * float64(time.Second)), true)
}

func GetCPUNice() (float64, error) {
	cpuStat, err := cpu.Times(false)
	if err != nil {
		return 0, err
	}
	if len(cpuStat) == 0 {
		return 0, nil
	}
	return cpuStat[0].Nice, nil
}

func GetCPUStealTime() (float64, error) {
	cpuStat, err := cpu.Times(false)
	if err != nil {
		return 0, err
	}
	if len(cpuStat) == 0 {
		return 0, nil
	}
	return cpuStat[0].Steal, nil
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

func GetCPUTemperature()(float64) {
	return 0.0
}

func GetCPUIdle() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return 100 - percentages[0], nil
}

func GetClocksPerSecond() (uint64) {
	return uint64(cpu.ClocksPerSec)
}
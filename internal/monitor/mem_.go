package monitor

import (
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryUsage() (float64, uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return vmStat.UsedPercent, vmStat.Used, nil
}

func GetMemoryTotal() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Total, nil
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

func GetMemoryFree() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Free, nil
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

func GetMemorySwapFree() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Free, nil
}

func GetMemoryTotalSwap() (uint64, error) {
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return 0, err
	}
	return swapStat.Total, nil
}

func GetMemoryTotalUsed() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Used, nil
}

func GetMemoryAvailable() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.Available, nil
}
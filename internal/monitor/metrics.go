package monitor

import (
	"runtime"
)

func GetSystemMetrics() map[string]float64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	metrics := map[string]float64{
		"cpu":    GetCPU_Usage(),     
		"memory": float64(memStats.Alloc) / 1024 / 1024,
	}
	return metrics
}

// CPU USAGE
func GetCPU_Usage() float64 {
	// Add logic
	return 0.0
}

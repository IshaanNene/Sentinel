package monitor

import (
	"runtime"
)

func GetSystemMetrics() map[string]float64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	metrics := map[string]float64{
		"cpu":    GetCPU_Usage(),      // Placeholder for actual CPU usage logic
		"memory": float64(memStats.Alloc) / 1024 / 1024,
	}
	return metrics
}

// Example placeholder for CPU usage logic
func GetCPU_Usage() float64 {
	// Add logic for CPU usage retrieval
	return 0.0
}

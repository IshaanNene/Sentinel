package monitor

import (
	_"log"
	"pkg/alerts"
)

type AlertConfig struct {
	CPUThreshold float64
	MemoryThreshold float64
}

func CheckAlerts(config AlertConfig, metrics map[string]float64) {
	if metrics["cpu"] > config.CPUThreshold {
		alerts.SendNotification("CPU usage exceeded threshold!")
	}
	if metrics["memory"] > config.MemoryThreshold {
		alerts.SendNotification("Memory usage exceeded threshold!")
	}
}

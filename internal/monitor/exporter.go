package monitor

import (
	"pkg/exporters"
	"/pkg/alerts"
)

func ExportMetricsToCSV(metrics map[string]float64, filename string) error {
	return exporters.ExportToCSV(metrics, filename)
}

func ExportMetricsToJSON(metrics map[string]float64, filename string) error {
	return exporters.ExportToJSON(metrics, filename)
}

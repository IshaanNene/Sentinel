package monitor

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"fmt"
)

func ExportMetricsToCSV(metrics map[string]float64, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for key, value := range metrics {
		if err := writer.Write([]string{key, fmt.Sprintf("%f", value)}); err != nil {
			return err
		}
	}
	return nil
}

func ExportMetricsToJSON(metrics map[string]float64, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(metrics)
}

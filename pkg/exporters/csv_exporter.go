package exporters

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ExportToCSV(metrics map[string]float64, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Metric", "Value"})
	for key, value := range metrics {
		writer.Write([]string{key, strconv.FormatFloat(value, 'f', 2, 64)})
	}
	return nil
}

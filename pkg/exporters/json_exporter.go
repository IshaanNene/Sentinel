package exporters

import (
	"encoding/json"
	"os"
)

func ExportToJSON(metrics map[string]float64, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(metrics)
}

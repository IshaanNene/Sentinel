package utils

import (
	"fmt"
	"os"
	"time"
)

func FormatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := unit, 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func LogToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	log := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), data)
	_, err = file.WriteString(log)
	return err
}

func FormatPercentage(value float64) string {
	return fmt.Sprintf("%.2f%%", value)
}

func CurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatMemory(total uint64, used uint64) string {
	return fmt.Sprintf("%s / %s", FormatBytes(used), FormatBytes(total))
}

func FormatDisk(total uint64, used uint64) string {
	return fmt.Sprintf("%s / %s", FormatBytes(used), FormatBytes(total))
}

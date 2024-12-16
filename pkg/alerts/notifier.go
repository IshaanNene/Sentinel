package alerts

import "log"

func SendNotification(message string) {
	log.Printf("ALERT: %s\n", message)
}

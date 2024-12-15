package alerts

import "log"

func SendNotification(message string) {
	log.Printf("ALERT: %s\n", message)
	// Add email/Slack/other notification logic here
}

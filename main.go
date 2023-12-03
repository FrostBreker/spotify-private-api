package spotifyprivateapi

import (
	"log"
)

func LogInfo(message string) {
	log.Printf("[INFO] --> %v", message)
}

func LogError(message string) {
	log.Printf("[ERROR] --> %v", message)
}

func LogDebug(message string) {
	log.Printf("[DEBUG] --> %v", message)
}

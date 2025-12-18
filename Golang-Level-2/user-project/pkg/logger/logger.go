package logger

import "log"

// Logger sederhana (siap dipakai nanti)
func Info(msg string) {
	log.Println("[INFO]", msg)
}

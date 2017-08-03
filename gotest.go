package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/hashicorp/logutils"
)

func main() {

	filter := &logutils.LevelFilter{
		Levels: []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer: os.Stderr,
	}
	log.SetOutput(filter)

	count := 1

	err := errors.New("Some application error")

	for {
		log.Printf("[DEBUG] Received GpsTransaction with key: %d", count)
		log.Printf("[INFO] Received GpsTransaction with key: %d", count)
		log.Printf("[WARN] Received GpsTransaction with key: %d", count)
		log.Printf("[ERROR] Received GpsTransaction with key: %d, err: %s", count, err)
		log.Printf("Received GpsTransaction with key: %d", count)
		count++
		time.Sleep(10 * time.Second)
	}

}
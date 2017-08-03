package main

import (
	"log"
	"time"
)

func main() {

	count := 1

	for {
		log.Printf("Received GpsTransaction with key: %d", count)
		count++
		time.Sleep(2 * time.Second)
	}

}
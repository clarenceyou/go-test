package main

import (
	"fmt"
	"time"
)

func main() {

	count := 1

	for {
		fmt.Printf("Testing go: %d\n", count)
		count++
		time.Sleep(2 * time.Second)
	}

}
package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Starting job that will fail continuously.")
	// Wait for 3 seconds before failing
	time.Sleep(3 * time.Second)
	log.Println("Encountered error, exiting with code 1.")
	os.Exit(1)
}

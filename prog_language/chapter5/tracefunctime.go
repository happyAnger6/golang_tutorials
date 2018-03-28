package main

import (
	"time"
	"log"
)

func trace(s string) func() {
	start := time.Now()
	log.Printf("enter %s", s)
	return func() {
		log.Printf("exit (%s) %s", s, time.Since(start))
	}
}

func main() {
	defer trace("main trace.")()

	time.Sleep(1*time.Second)
}

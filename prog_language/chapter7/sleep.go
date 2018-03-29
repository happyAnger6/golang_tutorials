package main

import (
	"flag"
	"time"
	"fmt"
)

func main() {
	var period = flag.Duration("period", 1 * time.Second, "sleep period.")
	flag.Parse()
	fmt.Printf("sleep for: %v period.\n", period)
	time.Sleep(*period)
	fmt.Println()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("timetick countdown.")

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("countdown: %v \n", countdown)
		<- tick
	}
}

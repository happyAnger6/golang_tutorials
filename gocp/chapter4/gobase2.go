package main

import (
	"fmt"
	 "time"
)

func main() {
	go func() {
		fmt.Println("goroutine.")
	}()

	time.Sleep(time.Millisecond)
}

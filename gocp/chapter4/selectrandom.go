package main

import "fmt"

func main() {
	chanCap := 5
	capChan := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case capChan <- 1:
		case capChan <- 2:
		case capChan <- 3:
		}
	}

	for i := 0; i < chanCap; i++ {
		fmt.Println(<-capChan)
	}

}

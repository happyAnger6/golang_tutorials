package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= 1500; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <- naturals
			if !ok {
				close(squares)
				break
			}
			squares <- x*x
		}
	}()

	for {
		if x, ok := <-squares; ok{
			fmt.Println(x)
		} else {
			break
		}
	}
}

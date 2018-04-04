package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i := 0; i <= 100; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Printf("write 2 chan:%d\r\n", i)
		}
	}
}

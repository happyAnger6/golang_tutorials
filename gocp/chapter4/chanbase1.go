package main

import (
	"fmt"
	"time"
)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	strChan := make(chan string, 2)
	go func() {
		<- syncChan1
		fmt.Println("received sync signal.")
		time.Sleep(time.Second)
		for {
			if elem, ok := <- strChan; ok {
				fmt.Printf("recevied [%s]\n", elem)
			}	else {
				fmt.Printf("strChan closed.")
				break
			}
		}
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _, elem := range []string{"a", "b", "c", "d", "e"} {
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Printf("send signal.\n")
				time.Sleep(time.Second)
			}
			strChan <- elem
			fmt.Printf("send str:%s\n", elem)
		}
		close(strChan)
		syncChan2 <- struct{}{}
	}()

	<- syncChan2
	<- syncChan2
	fmt.Println("done.")
}

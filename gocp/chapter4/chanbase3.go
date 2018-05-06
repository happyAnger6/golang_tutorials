package main

import (
	"fmt"
)

func receiver(strChan <-chan string,
	syncChan1 <-chan struct{},
	syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("receive signal.")
	for elem := range strChan {
		fmt.Println("Received", elem, "[receiver].")
	}
	fmt.Println("Stopped.[receiver].")
	syncChan2 <- struct{}{}
}

func sender(strChan chan <- string,
	syncChan1 chan <- struct{},
	syncChan2 chan <- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		if  elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("send a signal.[sender]")
		}
	}
	close(strChan)
	syncChan2 <- struct{}{}
}

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	strChan := make(chan string, 3)

	go receiver(strChan, syncChan1, syncChan2)
	go sender(strChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
	fmt.Println("done.")
}

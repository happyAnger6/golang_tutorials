package main

import (
	"sync"
	"runtime"
	"fmt"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		for {
			if elem, ok :=  <- mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		wait.Done()
	}()

	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			runtime.Gosched()
			fmt.Printf("countMap:%v\n", countMap)
		}
		close(mapChan)
		wait.Done()
	}()

	wait.Wait()
}

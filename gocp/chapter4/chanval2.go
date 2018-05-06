package main

import (
	"sync"
	"runtime"
	"fmt"
)

type Counter struct {
	count int
}

func (counter *Counter) String() string {
	return fmt.Sprintf("{count:%d}", counter.count)
}

var mapChan = make(chan map[string]*Counter, 1)

func main() {
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		for {
			if elem, ok :=  <- mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		wait.Done()
	}()

	go func() {
		countMap := map[string]*Counter{"count":&Counter{}}
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

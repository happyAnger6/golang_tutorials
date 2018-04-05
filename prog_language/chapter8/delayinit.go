package main

import (
	"fmt"
	"sync"
)

var res int
var one sync.Once

func initRes() {
	res++
}

func getRes() int {
	if res == 0 {
		one.Do(initRes)
		return res
	}
	return res
}

func main() {
	var wait sync.WaitGroup
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			fmt.Println(getRes())
			wait.Done()
		}()
	}

	wait.Wait()
}

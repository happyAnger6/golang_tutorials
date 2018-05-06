package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"an", "bob", "cidy", "dany", "eric"}
	for _, name := range names {
		go func(who string) {
			fmt.Println(who)
		}(name)
	}

	time.Sleep(time.Millisecond)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	items := []string{"a", "b"}
	fmt.Println(strings.Join(items, " "))
}

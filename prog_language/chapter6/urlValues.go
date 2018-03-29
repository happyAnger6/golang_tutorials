package main

import (
	"net/url"
	"fmt"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m)

	m = nil
	//m.Add("item", "3")
}

package main

import (
	"encoding/xml"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok, err := dec.Token()
	}
}

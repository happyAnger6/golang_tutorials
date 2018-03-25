package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, con); err != nil {
		log.Fatal(err)
	}
}

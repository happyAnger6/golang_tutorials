package main

import (
	"io"
)

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

func main() {
	var iw io.Writer
	f(iw)
}

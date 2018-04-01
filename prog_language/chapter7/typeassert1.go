package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	//c := w.(*bytes.Buffer)

	rw := w.(io.ReadWriter)
	fmt.Println(f)
	fmt.Printf("%v %#[1]T\n", rw)
}

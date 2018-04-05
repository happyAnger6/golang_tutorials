package main

import (
	"io"
	"image"
	"fmt"
	"os"
	"image/jpeg"
)

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Println(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality:95})
}

func main() {

}

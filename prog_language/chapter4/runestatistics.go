package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

func main() {
	rmaps := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int

	invalidCount := 0
	input := bufio.NewReader(os.Stdin)
	for {
		r, n, err := input.ReadRune()
		fmt.Printf("r:%q n:%d err:%v\n", r, n, err)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalidCount++
			continue
		}

		rmaps[r]++
		utflen[n]++
		fmt.Println(rmaps, utflen)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range rmaps {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcout\n")
	for i, n := range utflen {
		fmt.Printf("%q\t%d\n", i, n)
	}

	if invalidCount > 0 {
		fmt.Printf("invalid count:%d\n", invalidCount)
	}
}

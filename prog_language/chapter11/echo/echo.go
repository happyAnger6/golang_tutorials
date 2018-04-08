package main

import (
	"flag"
	"io"
	"os"
	"fmt"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "seperator")
)

var out io.Writer = os.Stdout

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprint(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

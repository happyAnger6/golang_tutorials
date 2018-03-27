package main

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
	"log"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	outline(nil, doc)
}

func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}

	for c := node.FirstChild; c != nil; c = node.NextSibling {
		outline(stack, c)
	}
}

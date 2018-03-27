package main

import (
	"golang.org/x/net/html"
	"os"
	"log"
	"fmt"
)


func visit(links []string, n *html.Node) []string {
	fmt.Printf("type:%v data:%v\n", n.Type, n.Data)
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	fmt.Printf("child:%v \n", n.FirstChild)
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		links = visit(links, c)
	}

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
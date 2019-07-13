package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)


func main() {
	var count = make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	countElement(count, doc)
	fmt.Printf("%s \t %s \n", "nodeName", "count")
	for name, num := range count {
		fmt.Printf("%s \t %d \n", name, num)
	}
}

func countElement(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		 countElement(count, c)
	}
}

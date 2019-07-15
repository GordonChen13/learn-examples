package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		err := outline(url)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(doc)
	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=%q", a.Key, a.Val)
		}
		if n.FirstChild != nil {
			fmt.Printf(">\n")
		} else {
			fmt.Printf(" />\n")
		}
		depth++
	} else if n.Type == html.TextNode {
		trimmed := strings.TrimSpace(n.Data)
		if trimmed != "" {
			fmt.Printf("%*s%s\n", depth*2, "", trimmed)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

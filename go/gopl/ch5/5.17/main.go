package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)
var depth int

func ElementsByTag(n *html.Node, tags ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	keep := make(map[string]bool, len(tags))
	for _, tag := range tags {
		keep[tag] = true
	}

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		_, ok := keep[n.Data]
		if ok {
			nodes = append(nodes, n)
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
		return true
	}
	post := func (n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		_, ok := keep[n.Data]
		if ok {
			if n.Type == html.ElementNode {
				depth--
				if n.FirstChild != nil {
					fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
				}
			}
		}
		return true
	}
	forEachNode(n, pre, post)
	return nodes
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool){
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

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: URL TAG...")
	}
	url := os.Args[1]
	tags := os.Args[2:]
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("get url content error: %s", err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("parse html error: %s", err)
	}
	ElementsByTag(doc, tags...)
	//for _, n := range nodes {
	//	fmt.Printf("%+v\n", n)
	//}
}

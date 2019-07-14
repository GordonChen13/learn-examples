package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "https://golang.org"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	fmt.Println("Parsing", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images := CountWordsAndImages(doc)
	fmt.Println("words:", words, "images", images)
}

func CountWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					images++
				}
			}
		}
	} else if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childwords, childimages := CountWordsAndImages(c)
		words += childwords
		images += childimages
	}

	return words, images
}

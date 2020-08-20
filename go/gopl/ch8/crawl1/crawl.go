package main

import (
	"fmt"
	"log"

	"github.com/GordonChen13/learn-examples/go/gopl/ch5/links"
)

func main() {
	worklist := make(chan []string)
	go func() {
		worklist <- []string{"http://hao123.com"}
	}()

	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			seen[link] = true
			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

package main

import (
	"fmt"
	"log"

	"github.com/GordonChen13/learn-examples/go/gopl/ch5/links"
)

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int
	n++
	go func() {
		worklist <- []string{"http://hao123.com"}
	}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			seen[link] = true
			n++
			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

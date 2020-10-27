package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/GordonChen13/learn-examples/go/gopl/ch5/links"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

func main() {
	flag.IntVar(&maxDepth, "d", 3, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	for _, link := range []string{"http://news.baidu.com", "http://music.qq.com"} {
		wg.Add(1)
		go crawl(link, 0, wg)
	}
	wg.Wait()
}

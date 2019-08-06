package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	*c += WordCounter(count)
	return count, nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	c := 1
	for _, b := range p {
		if b == '\n' {
			c++
		}
	}
	*lc += LineCounter(c)
	return c, nil
}

func main() {
	var wc WordCounter
	var lc LineCounter
	bytes, err := ioutil.ReadFile("words.txt")
	fmt.Println(string(bytes))
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	fmt.Fprintln(&wc, string(bytes))
	fmt.Fprintln(&lc, string(bytes))
	//wc.Write(bytes)
	//lc.Write(bytes)

	fmt.Println(wc)
	fmt.Println(lc)
}

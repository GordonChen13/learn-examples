package main

import (
	"bufio"
	"fmt"
    "log"
	"os"
)

func main() {
	files := os.Args[1:]
	var counts []map[string]int
	if len(files) == 0 {
        log.Fatal("you need to input a file")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			count := countLines(f)
			counts = append(counts, count)
			f.Close()
		}
	}

	for i, count := range counts {
		for _, n := range count {
			if n > 1 {
				fmt.Printf("%d\t %s\n", n ,files[i])
			}
		}
	}
}

func countLines(f *os.File) (counts map[string]int) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	return
}

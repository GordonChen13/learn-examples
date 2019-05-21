package main

import (
	"bufio"
	"fmt"
	"os"
)

var files []string = os.Args[1:]
var linesStdin = map[string]int{}
var counts = map[string]map[string]int{}

func main() {
	if len(files) == 0 {
		countStdin(os.Stdin, linesStdin)
		for line, n := range linesStdin {
			if n > 1 {
				fmt.Printf("%s\t%d\n", line, n)
			}
		}
	} else {
		for _, arg := range files {
			counts[arg] = map[string]int{}
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for fileName, innerMap := range counts {
			for line, n := range innerMap {
				if n > 1 {
					fmt.Printf("%s\t%s\t%d\n", fileName, line, n)
				}
			}
		}
	}
}

func countStdin(f *os.File, linesStdin map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		linesStdin[input.Text()]++
	}

}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}

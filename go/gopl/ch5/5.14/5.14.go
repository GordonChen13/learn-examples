package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":      {"discrete math"},
	"databases":            {"data structures"},
	"discrete math":        {"intro to programing"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programing languages": {"data structures", "computer organization"},
}

func main() {
	var course string
	for course = range prereqs { // 获取随机字段
		break
	}
	breadthFirst(deps, []string{course})
}

func deps(coures string) []string {
	fmt.Println(coures)
	return prereqs[coures]
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

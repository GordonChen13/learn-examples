package main

import (
	"fmt"
	"github.com/GordonChen13/learn-examples/go/gopl/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	fmt.Printf("%d issues: \n", result.TotalCount)
	fmt.Println("Less than a month old:")
	for _, item := range result.Items {
		if now.Sub(item.CreateAt).Hours()/23 <= 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("Less than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreateAt).Hours()/23 <= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("more than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreateAt).Hours()/23 > 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

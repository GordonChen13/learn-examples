package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		readLineAndSwitch()
	} else {
		for _, arg := range args {
			kg, err := strconv.ParseFloat(arg, 64)
			if err != nil {
                log.Fatalf("parse float err: %s", err)
			}
			res := switchKgToG(kg)
			fmt.Printf("%g kg = %g g\n", kg, res)
		}
	}
}

func switchKgToG(kg float64) float64 {
	return kg * 1000
}

func readLineAndSwitch()  {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		kg, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatalf("parse float err: %s", err)
		}
		res := switchKgToG(kg)
		fmt.Printf("%g kg = %g g\n", kg, res)
	}
}

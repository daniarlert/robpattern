package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: robpatter <re>")
	}

	matches := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if Match(os.Args[1], line) {
			fmt.Fprintln(os.Stdout, line)
			matches++
		}
	}

	if matches == 0 {
		fmt.Fprintln(os.Stdout, "no matches found")
	}
}

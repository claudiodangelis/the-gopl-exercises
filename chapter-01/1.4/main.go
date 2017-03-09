package main

import (
	"fmt"
	"os"
	"strings"
)

import "bufio"

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		counts[file] = make(map[string]int)
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
			continue
		}
		countLines(file, f, counts)
		f.Close()
	}
	sep := ""
	for name, hash := range counts {
		l, _ := fmt.Printf("%sDupes in %s\n", sep, name)
		sep = "\n"
		fmt.Printf("%s\n", strings.Repeat("-", l-1))
		for line, count := range hash {
			if count > 1 {
				fmt.Printf("%s\t%d\n", line, count)
			}
		}
	}
}

func countLines(name string, f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[name][input.Text()]++
	}
}

// Modify dup2 to print the names of all files in which each duplicate line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileLines := make(map[string][]string)  // {"filename": []lines}
	var arg string

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileLines, arg)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex4: %v\n", err)
				continue
			}
			countLines(f, counts, fileLines, arg)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for filename, lines := range fileLines {
				for _, l := range lines {
					if l == line {
						fmt.Printf("%s\n", filename)
						break
					}
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileLines map[string][]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "EOF" {
			break
		}
		counts[input.Text()]++
		fileLines[arg] = append(fileLines[arg], input.Text())
	}
}

package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	for input.Scan() {
		line := input.Text()
		counts[line] = counts[line] + 1
		// counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

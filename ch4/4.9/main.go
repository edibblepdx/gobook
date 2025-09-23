// Exercise 4.9: Write a program wordfreq to report the frequency of each word
// in an input text file. Call input.Split(bufio.ScanWords) before the first call
// to Scan to break the input into words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v", err)
		os.Exit(1)
	}

	fmt.Print("\ncount\tword\n")
	for k, v := range counts {
		fmt.Printf("%d\t%q\n", v, k)
	}
}

// Exercise 1.4: Modify dup2 to print the names of all files in which each
// duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineInfo struct {
	count int
	files map[string]bool
}

func main() {
	lines := make(map[string]lineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", lines)
	} else {
		for _, fname := range files {
			file, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(file, fname, lines)
			file.Close()
		}
	}
	for line, info := range lines {
		if info.count > 1 {
			fmt.Printf("%d\t%s\n", info.count, line)
			for fname := range info.files {
				fmt.Printf("\t> %s\n", fname)
			}
		}
	}
}

func countLines(file *os.File, fname string, lines map[string]lineInfo) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		entry, ok := lines[input.Text()]
		if !ok {
			entry = lineInfo{0, make(map[string]bool)}
		}
		entry.count++
		entry.files[fname] = true
		lines[input.Text()] = entry
	}
	// NOTE: ignoring potential errors from input.Err()
}

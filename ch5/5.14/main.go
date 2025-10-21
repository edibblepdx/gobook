// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 5.14: Use the breadthFirst function to explore a different structure.
// I'll make a simple find utility with the -name option provided.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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

func find(path string) (list []string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Print(err)
	}
	for _, entry := range dir {
		if entry.Name() == name {
			fmt.Println(filepath.Join(path, entry.Name()))
		}
		if entry.IsDir() {
			list = append(list, filepath.Join(path, entry.Name()))
		}
	}
	return list
}

var (
	start string
	name  string
)

func init() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: run <start> <name>")
	}

	start = os.Args[1]
	name = os.Args[2]
}

func main() {
	breadthFirst(find, []string{start})
}

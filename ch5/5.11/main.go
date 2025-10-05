// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 5.11: The instructor of the linear algeabra course decides that
// calculus is now a prerequisite. Extend the topoSort function to report cycles.

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	//"linear algebra": {"calculus"},
	"intro to programming": {"compilers"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

// !+main
func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func filterKey[K comparable, V any](m map[K]V, f func(K, V) bool) (result []K) {
	for k, v := range m {
		if f(k, v) {
			result = append(result, k)
		}
	}
	return
}

func topoSort(m map[string][]string) (order []string, err error) {
	var visitAll func(items []string)

	// True when all children are sorted
	sorted := make(map[string]bool)

	visitAll = func(items []string) {
		for _, item := range items {
			if err != nil {
				return
			}
			done, seen := sorted[item]
			if !done && seen {
				err = fmt.Errorf("cycle: %v", strings.Join(filterKey(sorted,
					func(_ string, v bool) bool { return !v }), " -> "))
				return
			}
			if !seen {
				sorted[item] = false
				visitAll(m[item])
				sorted[item] = true
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order, err
}

//!-main

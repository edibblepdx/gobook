// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 5.10: rewrite topoSort to use maps instead of slices and eliminate
// the initial sort. Verify that the results, though nondeterministic, are valid
// topological orderings.

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

type graph map[string]edges
type edges map[string]bool

// prereqs maps computer science courses to their prerequisites.
var prereqs = graph{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":  {"discrete math": true},
	"databases":        {"data structures": true},
	"discrete math":    {"intro to programming": true},
	"formal languages": {"discrete math": true},
	"networks":         {"operating systems": true},

	"operating systems": {
		"data structures":       true,
		"computer organization": true,
	},

	"programming languages": {
		"data structures":       true,
		"computer organization": true,
	},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(g graph) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(e edges)

	visitAll = func(e edges) {
		for node := range e {
			if !seen[node] {
				seen[node] = true
				visitAll(g[node])
				order = append(order, node)
			}
		}
	}

	sg := make(graph)
	sg["o"] = make(edges)
	for node := range g {
		sg["o"][node] = true
	}

	visitAll(sg["o"])
	return order
}

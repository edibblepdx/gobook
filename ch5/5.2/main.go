// Exercise 5.2: Write a function to populate a mapping from element names--p,
// div, span, and so on--to the number of elements with that name in an HTML
// document tree.

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for k, v := range countelems(make(map[string]int), doc) {
		fmt.Printf("%d\t%s\n", v, k)
	}
}

func countelems(e map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		e[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		e = countelems(e, c)
	}
	return e
}

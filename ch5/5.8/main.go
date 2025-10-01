// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 5.8: Modify forEachNode so that the pre and post functions return a
// boolean result indicating whether to continue the traversal. Use it to write
// a function ElementByID with the following signature that finds the first HTML
// element with the specified id attribute. The function should stop the
// traversal as soon as a match a found.
//
// func ElementByID(doc *html.Node, id string) *html.Node

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: run <url> <id>")
		os.Exit(1)
	}
	n, err := ElementByID(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(n.Data, n.Attr)
}

func ElementByID(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return elementByID(doc, id), nil
}

func elementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node
	forEachNode(doc, func(n *html.Node) (stop bool) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					node = n
					return true
				}
			}
		}
		return false
	}, nil)

	return node
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) (stop bool)) (stop bool) {
	if pre != nil {
		if pre(n) {
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(c, pre, post) {
			return true
		}
	}

	if post != nil {
		if post(n) {
			return true
		}
	}

	return false
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 5.7: Develop startElement and endElement in toa general HTML pretty-
// printer.

// I tested it on my website and it looks good.
// Worth coming back to after chapter 11.

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

// !+startend
var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		depth++
		for _, a := range n.Attr {
			fmt.Printf(" %s=%s", a.Key, a.Val)
		}
		if n.FirstChild != nil {
			fmt.Println(">")
		} else {
			fmt.Println("/>")
		}
	case html.TextNode:
		lines := strings.Split(n.Data, "\n")
		for _, line := range lines {
			t := strings.TrimSpace(line)
			if t != "" {
				fmt.Printf("%*s  %s\n", depth*2, "", t)
			}
		}
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend

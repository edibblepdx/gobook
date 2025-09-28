// Exercise 5.3: Write a function to print the contents of all text nodes in an
// HTML document tree. Do not descend into <script> or <style> elements, since
// their contents are not visible in a web browser.

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	text(doc)
}

func text(n *html.Node) {
	if n.Type == html.TextNode {
		if t := strings.TrimSpace(n.Data); t != "" {
			fmt.Println(t)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode &&
			(c.Data == "script" || c.Data == "style") {
			continue
		}
		text(c)
	}
}

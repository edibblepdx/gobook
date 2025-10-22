// Exercise 5.17: Write a variadic function ElementsByTagName that, given an HTML
// node tree and zero or more names, returns all of the elements that match those
// one of those names.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	"golang.org/x/net/html"
)

func ElementsByTagName(n *html.Node, tags ...string) []*html.Node {
	var elems []*html.Node
	var collect func(n *html.Node)

	collect = func(n *html.Node) {
		if n.Type == html.ElementNode && slices.Contains(tags, n.Data) {
			elems = append(elems, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			collect(c)
		}
	}
	collect(n)

	return elems
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	fmt.Println("images: ", len(images))
	fmt.Println("headings: ", len(headings))
}

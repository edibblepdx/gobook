// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 5.13: Modify crawl to make local copies of the pages it finds,
// creating directories as necessary. Don't make copies of pages that come from
// a different domain.

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopl.io/ch5/links"
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

func crawl(url string) []string {
	domain := strings.Split(url, "/")[2]

	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	n := 0
	for _, val := range list {
		// including subdomains
		if strings.Contains(val, domain) && !strings.ContainsAny(val, "#?=") {
			savePage(url)
			list[n] = val
			n++
		}
	}

	return list[:n]
}

var prefix = regexp.MustCompile(`^https?://`)

func savePage(url string) error {
	path := prefix.ReplaceAllString(url, "")
	// some pages might already have the .html suffix in path
	path, _ = strings.CutSuffix(path, ".html")
	path, _ = strings.CutSuffix(path, "/")
	path = "data/" + path + ".html"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		return err
	}

	return nil
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}

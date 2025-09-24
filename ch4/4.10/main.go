// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.10: Modify issues to report the results in age categories, say less
// than a month old, less than a year old, and more that a year old.

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const ( // Durations in nanoseconds
	MONTH = 30 * 24 * time.Hour
	YEAR  = 365 * 24 * time.Hour
	MAX   = 1<<63 - 1 // 290 years
)

func filter(issues []*github.Issue, min, max time.Duration) (filtered []*github.Issue) {
	for _, item := range issues {
		since := time.Since(item.CreatedAt)
		if min < since && since <= max {
			filtered = append(filtered, item)
		}
	}
	return
}

func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	var filtered []*github.Issue

	filtered = filter(result.Items, -1, MONTH)
	fmt.Printf("\n%d <= month:\n", len(filtered))
	printIssues(filtered)

	filtered = filter(result.Items, MONTH, YEAR)
	fmt.Printf("\n%d <= year:\n", len(filtered))
	printIssues(filtered)

	filtered = filter(result.Items, YEAR, MAX)
	fmt.Printf("\n%d > year:\n", len(filtered))
	printIssues(filtered)
}

//!-

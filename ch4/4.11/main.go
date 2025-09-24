// Exercise 4.11: Build a tool that lets users create, read, update, and close
// GitHub issues from the command line, invoking their preferred text editor when
// substantial text input is required.

// Issues API:
// https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28

// To use, create a fine-grained personal access token through GitHub developer
// settings and set the environment variable GH_PAT to the value of your token.
// Give the token issues permissions and a short expiration date.
// Read more here:
// https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token

package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	//"net/url"
	"fmt"
	//"gopl.io/ch4/github"
)

var (
	PAT    string // personal access token
	URL    string
	METHOD string
)

var method = map[string]func(){
	"create": func() {},
	"read":   func() {},
	"update": func() {},
	"close":  func() {},
}

func init() {
	PAT = os.Getenv("GH_PAT")
	if PAT == "" {
		log.Fatal("no personal access token in GH_PAT")
	}

	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("usage: <url> <method>")
	}
	URL = args[0]
	if https := strings.Contains(URL, "https://"); !https {
		URL = "https://" + URL
	}
	METHOD = args[1]
	if _, ok := method[METHOD]; !ok {
		log.Fatalf(
			"invalid method: %s\nvalid methods: create, read, update, delete",
			METHOD,
		)
	}
}

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Print(resp.Body)
}

// Exercise 4.11: Build a tool that lets users create, read, update, and close
// GitHub issues from the command line, invoking their preferred text editor when
// substantial text input is required.

// Issues API:
// https://docs.github.com/en/rest/issues?apiVersion=2022-11-28
// Issues PAT API:
// https://docs.github.com/en/rest/authentication/endpoints-available-for-fine-grained-personal-access-tokens?apiVersion=2022-11-28#issues
// REST API endpoints for PAT:
// https://docs.github.com/en/rest/orgs/personal-access-tokens?apiVersion=2022-11-28

// To use, create a fine-grained personal access token through GitHub developer
// settings and set the environment variable GH_PAT to the value of your token.
// Give the token issues permissions and a short expiration date.
// Read more here:
// https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"gobook/ch4/4.11/github"
)

var (
	PAT          string // personal access token
	METHOD       string
	OWNER        string
	REPO         string
	ISSUE_NUMBER string
)

func init() {
	PAT = os.Getenv("GH_PAT")
	if PAT == "" {
		log.Fatal("no personal access token in GH_PAT")
	}

	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("usage: <method> <url> [issue_number]")
	}
	METHOD = args[0]
	if _, ok := do[METHOD]; !ok {
		log.Fatalf(
			"invalid method: %s\nvalid methods: create, read, update, delete",
			METHOD)
	}
	OWNER, REPO = parseUrl(args[1])

	// (optional) issue number
	if len(args) > 2 {
		ISSUE_NUMBER = args[2]
	}
}

func parseUrl(s string) (string, string) {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatalf("invalid url: %v", err)
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		log.Fatal("invalid repo")
	}
	return parts[0], parts[1]
}

func main() {
	if err := do[METHOD](http.DefaultClient); err != nil {
		log.Fatal(err)
	}
}

func setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+PAT)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
}

func updateIssue(old string) (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	title = strings.TrimSpace(title)

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	f, err := os.CreateTemp("", "issue-*.md")
	if err != nil {
		return "", "", err
	}
	defer os.Remove(f.Name())

	if _, err := f.WriteString(old); err != nil {
		return "", "", err
	}
	f.Close()

	cmd := exec.Command(editor, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", "", err
	}

	body, err := os.ReadFile(f.Name())
	if err != nil {
		return "", "", err
	}

	return title, string(body), nil
}

var do = map[string]func(*http.Client) error{
	"create": func(c *http.Client) error {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues",
			OWNER, REPO)

		issueTitle, issueBody, err := updateIssue("Issue Body")
		if err != nil {
			return err
		}

		body := map[string]string{
			"title": issueTitle,
			"body":  issueBody,
		}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
		if err != nil {
			return err
		}
		setHeaders(req)

		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return nil
	},
	"read": func(c *http.Client) error {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues",
			OWNER, REPO)

		if ISSUE_NUMBER != "" {
			url = url + "/" + ISSUE_NUMBER
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		setHeaders(req)

		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("search query failed: %s", resp.Status)
		}

		if ISSUE_NUMBER == "" {
			var result []github.Issue
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				return err
			}
			for _, item := range result {
				item.Mini()
			}
		} else {
			var result github.Issue
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				return err
			}
			result.Show()
		}

		return nil
	},
	"update": func(c *http.Client) error {
		if ISSUE_NUMBER == "" {
			return fmt.Errorf("require issue number")
		}

		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s",
			OWNER, REPO, ISSUE_NUMBER)

		// Old Issue

		_req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		setHeaders(_req)

		_resp, err := c.Do(_req)
		if err != nil {
			return err
		}
		defer _resp.Body.Close()

		var old github.Issue
		if err := json.NewDecoder(_resp.Body).Decode(&old); err != nil {
			return err
		}

		// New Issue

		issueTitle, issueBody, err := updateIssue(old.Body)
		if err != nil {
			return err
		}

		body := map[string]string{
			"title": issueTitle,
			"body":  issueBody,
		}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("PATCH", url, bytes.NewReader(jsonBody))
		if err != nil {
			return err
		}
		setHeaders(req)

		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("search query failed: %s", resp.Status)
		}

		return nil
	},
	"close": func(c *http.Client) error {
		if ISSUE_NUMBER == "" {
			return fmt.Errorf("require issue number")
		}

		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s",
			OWNER, REPO, ISSUE_NUMBER)

		body := map[string]string{"state": "closed"}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("PATCH", url, bytes.NewReader(jsonBody))
		if err != nil {
			return err
		}
		setHeaders(req)

		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("search query failed: %s", resp.Status)
		}

		return nil
	},
}

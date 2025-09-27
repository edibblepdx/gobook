package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Issue struct {
	Id        int64
	Number    int
	HtmlUrl   string `json:"html_url"`
	State     string
	Title     string
	Body      string
	User      *User
	Milestone *Milestone
	Labels    []*Label
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
}

type User struct {
	Id      int64
	Login   string
	HtmlUrl string `json:"html_url"`
}

type Milestone struct {
	Id           int64
	Number       int
	HtmlUrl      string `json:"html_url"`
	State        string
	Title        string
	Description  string
	Creator      *User
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     time.Time `json:"closed_at"`
}

type Label struct {
	Name string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	const IssuesURL = "https://api.github.com/search/issues"

	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?per_page=100&q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

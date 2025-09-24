package github

import (
	"fmt"
	"time"
)

type Issue struct {
	Number    int
	HtmlUrl   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	Labels    []*Label
}

type User struct {
	Login   string
	HtmlUrl string `json:"html_url"`
}

type Label struct {
	Name string
}

func (item *Issue) Mini() {
	fmt.Printf("#%-5d %9.9s %.55s\n",
		item.Number, item.User.Login, item.Title)
}

func (item *Issue) Show() {
	fmt.Printf("#%d %s: %s\n%s\n[%s]\n",
		item.Number,
		item.User.Login,
		item.Title,
		item.Body,
		item.State)
}

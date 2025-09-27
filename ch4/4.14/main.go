// Exercise 4.14: Create a web server that queries GitHub once and then allows
// navigation of the list of bug reports, milestones, and users.

// Queries are paginated and return 30 results by default with a max of 100.
// You can't get any more in a single query.

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"gobook/ch4/4.14/github"
)

var (
	IssuesResult *github.IssuesSearchResult
	Users        map[int64]*github.User
	Milestones   map[int64]*github.Milestone
)

func init() {
	Users = make(map[int64]*github.User)
	Milestones = make(map[int64]*github.Milestone)
}

func main() {
	var err error
	IssuesResult, err = github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range IssuesResult.Items {
		if item.User != nil {
			Users[item.User.Id] = item.User
		}
		if item.Milestone != nil {
			Milestones[item.Milestone.Id] = item.Milestone
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/issues", http.StatusMovedPermanently)
	})
	http.HandleFunc("/issues", issues)
	http.HandleFunc("/milestones", milestones)
	http.HandleFunc("/users", users)

	log.Print("running on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var IssueList = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(`
<h1>{{.TotalCount}} issues</h1>
<a href="/issues">issues</a>
<a href="/users">users</a>
<a href="/milestones">milestones</a>
<br /><br />
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>Age</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HtmlUrl}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HtmlUrl}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HtmlUrl}}'>{{.Title}}</a></td>
  <td>{{.CreatedAt | daysAgo}} days</td>
</tr>
{{end}}
</table>
`))

func issues(w http.ResponseWriter, r *http.Request) {
	if err := IssueList.Execute(w, IssuesResult); err != nil {
		log.Print(err)
	}
}

var UserList = template.Must(template.New("userlist").Parse(`
<h1>users</h1>
<a href="/issues">issues</a>
<a href="/users">users</a>
<a href="/milestones">milestones</a>
<br /><br />
<table>
<tr style='text-align: left'>
  <th>Name</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HtmlUrl}}'>{{.Login}}</a></td>
</tr>
{{end}}
</table>
`))

func users(w http.ResponseWriter, r *http.Request) {
	if err := UserList.Execute(w, Users); err != nil {
		log.Print(err)
	}
}

var MilestoneList = template.Must(template.New("milestonelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(`
<h1>milestones</h1>
<a href="/issues">issues</a>
<a href="/users">users</a>
<a href="/milestones">milestones</a>
<br /><br />
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Open Issues</th>
  <th>Closed Issues</th>
  <th>Creator</th>
  <th>Title</th>
  <th>Age</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HtmlUrl}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td>{{.OpenIssues}}</td>
  <td>{{.ClosedIssues}}</td>
  <td><a href='{{.Creator.HtmlUrl}}'>{{.Creator.Login}}</a></td>
  <td><a href='{{.HtmlUrl}}'>{{.Title}}</a></td>
  <td>{{.CreatedAt | daysAgo}} days</td>
</tr>
{{end}}
</table>
`))

func milestones(w http.ResponseWriter, r *http.Request) {
	if err := MilestoneList.Execute(w, Milestones); err != nil {
		log.Print(err)
	}
}

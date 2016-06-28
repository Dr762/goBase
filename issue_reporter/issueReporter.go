package github

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

const templ = `{{.TotalCount}} issues:
{{range.Items}}--------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

var issueList = template.Must(template.New("issueList").Parse(
	`<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	</tr>
    {{range.Items}}
	<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
    `))

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Mode string

const (
	HTMLMode    Mode = "html_mode"
	ConsoleMode      = "console_mode"
)

func SearchIssues(terms []string, mode Mode) IssuesSearchResult {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	switch mode {
	case ConsoleMode:
		if err := report.Execute(os.Stdout, result); err != nil {
			return err
		}
	case HTMLMode:
		if err := issueList.Execute(os.Stdout, result); err != nil {
			return err
		}
	}

	return nil

}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

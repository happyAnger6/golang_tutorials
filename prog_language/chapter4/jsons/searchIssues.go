package main

import (
	"time"
	"net/url"
	"strings"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"text/template"
	"os"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, ""))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil;
}

const templ = `{{.TotalCount}} issues:
{{ range .Items}}
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s" }}
Age: {{.CreatedAt | daysAgo }} days
{{ end }}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	issueSearchResults, err := SearchIssues([]string{"url, login"})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range issueSearchResults.Items {
		fmt.Printf("%#v \n", item)
	}

	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(os.Stdout, issueSearchResults)


}

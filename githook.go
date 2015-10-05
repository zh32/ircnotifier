package ircnotifier

import (
	"net/http"
	"encoding/json"
	"strings"
	"io"
	"fmt"
)

type GitHook struct {}

func (githook GitHook) Execute(request *http.Request) *[]string {
	event, err := githook.parseRequestBody(request.Body)
	if err != nil {
		return nil
	}

	var lines []string
	lines = make([]string, 6)
	lines[0] = fmt.Sprintf(
		"%s pushed \u000310%d\u000f commit(s) to \u0002%s/%s",
		event.UserName,
		event.TotalCommitCount,
		event.Repository.Name,
		strings.Split(event.Ref, "/")[2])

	for i, commit := range event.Commits {
		lines[i+1] = fmt.Sprintf(
			"    %s \u000310%s\u000f: %s",
			commit.Id[:7],
			commit.Author.Name,
			commit.Message)
	}
	return &lines
}

func (githook *GitHook) parseRequestBody(reader io.Reader) (*GitEvent, error) {
	var pushEvent GitEvent
	decoder := json.NewDecoder(reader)
	return &pushEvent, decoder.Decode(&pushEvent)
}

type GitEvent struct {
	Repository       Repository `json:"repository"`
	Ref              string `json:"ref"`
	UserName         string `json:"user_name"`
	Commits          []Commit `json:"commits"`
	TotalCommitCount int `json:"total_commits_count"`
}

type Repository struct {
	Name string
	Url  string
}

type Commit struct {
	Id      string
	Message string
	Author  Author
}

type Author struct {
	Name  string
	Email string
}
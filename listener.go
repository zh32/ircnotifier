package ircnotifier

import (
	"net/http"
	"encoding/json"
	"strings"
	"io"
)

type NotificationListener struct {
	Writer NotificationWriter
	Key *string
	Host *string
}
func (listener *NotificationListener) Listen() {
	http.HandleFunc("/notify", func (response http.ResponseWriter, request *http.Request) {
		if !(listener.checkAccess(request)) {
			response.WriteHeader(http.StatusForbidden)
			return
		}
		pushEvent, err := listener.parseRequestBody(request.Body)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		listener.Writer.Write(
			&Notification{
				pushEvent,
				strings.Split(request.URL.Query().Get("c"), ","),
			},
		)
		response.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(*listener.Host, nil)
}

func (listener *NotificationListener) checkAccess(request *http.Request) bool {
	return *listener.Key == request.URL.Query().Get("k")
}

func (listener *NotificationListener) parseRequestBody(reader io.Reader) (*PushEvent, error) {
	var pushEvent PushEvent
	decoder := json.NewDecoder(reader)
	return &pushEvent, decoder.Decode(&pushEvent)
}

type Notification struct {
	PushEvent *PushEvent
	Channels []string
}

type PushEvent struct {
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
package ircnotifier

import (
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)

type JenkinsHook struct {}

func (jenkinshook JenkinsHook) Execute(request *http.Request) *[]string {
	event, err := jenkinshook.parseRequestBody(request.Body)
	if err != nil {
		return nil
	}
	message := jenkinshook.getLine(event)
	if message == nil {
		return nil
	}
	var lines []string
	lines = make([]string, 1)
	lines[0] = *message
	return &lines
}

func (jenkinshook JenkinsHook) getLine(event *JenkinsEvent) *string {
	var line string
	switch event.Build.Phase {
	case "STARTED":
		line = fmt.Sprintf(
			"%s: %s build #%d: %s",
			event.Name,
			event.Build.Phase,
			event.Build.Number,
			event.Build.FullUrl,
		)
	case "FINISHED":
		line = fmt.Sprintf(
			"%s: %s build #%d with status: %s%s",
			event.Name,
			event.Build.Phase,
			event.Build.Number,
			jenkinshook.statusColor(event.Build.Status),
			event.Build.Status,
		)
	}
	return &line;
}

func (jenkinshook JenkinsHook) statusColor(status string) string {
	switch status {
	case "SUCCESS":
		return "\u00039"
	default:
		return "\u00034"
	}
}
func (jenkinshook *JenkinsHook) parseRequestBody(reader io.Reader) (*JenkinsEvent, error) {
	var jenkinsEvent JenkinsEvent
	decoder := json.NewDecoder(reader)
	return &jenkinsEvent, decoder.Decode(&jenkinsEvent)
}

type JenkinsEvent struct {
	Name string
	Url string
	Build Build
}

type Build struct {
	FullUrl string `json:"full_url"`
	Phase string
	Status string
	Url string
	Number int
}

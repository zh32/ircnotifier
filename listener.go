package ircnotifier

import (
	"net/http"
	"strings"
)

var instance NotificationListener

type NotificationListener struct {
	Writer *NotificationWriter
	Key *string
	Host *string
	hooks map[string]Hook
}

func NewListener(writer *NotificationWriter, key, host *string) *NotificationListener {
	instance = NotificationListener{Writer: writer, Key: key, Host: host}
	return &instance
}

func (listener *NotificationListener) Listen() {
	listener.register("/git", &GitHook{})
	listener.register("/jenkins", &JenkinsHook{})
	http.ListenAndServe(*listener.Host, nil)
}

func (listener *NotificationListener) register(path string, hook Hook) {
	if listener.hooks == nil {
		listener.hooks = make(map[string]Hook)
	}
	listener.hooks[path] = hook
	http.HandleFunc(path, handleL)
}

func handleL(response http.ResponseWriter, request *http.Request) {
	if !(instance.checkAccess(request)) {
		response.WriteHeader(http.StatusForbidden)
		return
	}
	hook := instance.hooks[request.URL.Path]
	lines := hook.Execute(request)
	if lines == nil {
		response.WriteHeader(http.StatusOK)
		return
	}
	channels := strings.Split(request.URL.Query().Get("c"), ",")
	instance.Writer.Write(&Notification{lines, &channels})
	response.WriteHeader(http.StatusOK)
}

func (listener *NotificationListener) checkAccess(request *http.Request) bool {
	return *listener.Key == request.URL.Query().Get("k")
}

type Notification struct {
	Messages *[]string
	Channels *[]string
}

type Hook interface {
	Execute(request *http.Request) *[]string
}
package ircnotifier

import (
	irc "github.com/fluffle/goirc/client"
	"fmt"
	"strings"
)

type NotificationWriter struct {
	Client *irc.Conn
}

func (writer *NotificationWriter) Write(notification *Notification) {
	for _, channel := range notification.Channels {
		writer.Client.Privmsg(channel, fmt.Sprintf(
			"%s pushed \u000310%d\u000f commit(s) to \u0002%s/%s",
			notification.PushEvent.UserName,
			notification.PushEvent.TotalCommitCount,
			notification.PushEvent.Repository.Name,
			strings.Split(notification.PushEvent.Ref, "/")[2]))

		for _, commit := range notification.PushEvent.Commits {
			writer.Client.Privmsg(
				channel,
				fmt.Sprintf(
					"    %s \u000310%s\u000f: %s",
					commit.Id[:7],
					commit.Author.Name,
					commit.Message))
		}
	}
}

package ircnotifier

import (
	irc "github.com/fluffle/goirc/client"
)

type NotificationWriter struct {
	Client *irc.Conn
}

func (writer *NotificationWriter) Write(notification *Notification) {
	for _, channel := range *notification.Channels {
		for _, message := range *notification.Messages {
			writer.Client.Privmsg(channel, message)
		}
	}
}

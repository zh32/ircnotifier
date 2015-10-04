package ircnotifier

import (
	"strings"
	irc "github.com/fluffle/goirc/client"
)

type Commands struct {
	commands map[string]Command
}

func (registry *Commands) Register(name string, command Command) {
	if registry.commands == nil {
		registry.commands = make(map[string]Command)
	}
	registry.commands[name] = command
}

func (registry *Commands) ParseCommand(line string) Command {
	if strings.Index(line, ".") == 0 {
		end := strings.Index(line, " ");
		if (end == -1) {
			end = len(line)
		}
		name := line[1:end]
		return registry.commands[name]
	}
	return nil
}

type Command interface {
	Execute(conn *irc.Conn, line *irc.Line)
}

type HelloWorld struct {}

func (command HelloWorld) Execute(conn *irc.Conn, line *irc.Line) {
	conn.Privmsg(line.Target(), "ich bin ein bot")
}
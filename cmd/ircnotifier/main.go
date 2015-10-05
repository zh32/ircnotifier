package main

import (
	"fmt"
	notifier "github.com/zh32/ircnotifier"
	irc "github.com/fluffle/goirc/client"
	"flag"
	"strings"
)

var quit chan bool
var registry notifier.Commands

func main() {
	key := flag.String("key", "secret", "secret key")
	host := flag.String("host", "localhost:8080", "listen address")
	ircServer := flag.String("server", "irc.freenode.net", "irc server")
	ircNick := flag.String("nick", "Notifier", "irc nick name")
	ircChannels := flag.String("channels", "##mcseu", "comma separated list of channels")
	flag.Parse()
	fmt.Println(*ircChannels)
	registry = notifier.Commands{}
	registry.Register("hello", &notifier.HelloWorld{})

	client := connectIrc(ircServer, ircNick, ircChannels)

	listener := notifier.NewListener(
		&notifier.NotificationWriter{Client: client},
		key,
		host,
	)
	listener.Listen()

	quit = make(chan bool)

	<- quit
}

func connectIrc(server *string, nick *string, ircChannels *string) *irc.Conn {
	client := irc.SimpleClient(*nick)
	client.HandleFunc(irc.CONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			for _, channel := range strings.Split(*ircChannels, ",") {
				conn.Join(channel)
			}
		})
	client.HandleFunc(irc.PRIVMSG, h_PRIVMSG)
	if err := client.ConnectTo(*server); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}
	return client;
}

func h_PRIVMSG(conn *irc.Conn, line *irc.Line) {
	command := registry.ParseCommand(line.Text())
	if command != nil {
		fmt.Printf("Executing command %T\n", command)
		command.Execute(conn, line)
	}
}
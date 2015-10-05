ircnotifier
===========

Small application that accepts http requests and relays its formatted payload to several irc channels.
Supported payloads:

* Github/Gitlab Push Events on ``/git``
* Jenkins Notification Plugin on ``/jenkins``

##Installation

``go get github.com/zh32/ircnotifier``
``go install github.com/zh32/ircnotifier/cmd/ircnotifier``

##Usage

####Run the program
``./ircnotifier [OPTION]...``

####Options
```
-channels string
    comma separated list of channels (default "##mcseu")
-host string
    listen address (default "localhost:8080")
-key string
    secret key (default "secret")
-nick string
    irc nick name (default "Notifier")
-server string
    irc server (default "irc.freenode.net")
```

####Add url to your services
To receive a notification you need to pass the secret and the channels with the request. For example:
``http://localhost:8080/jenkins?k=mysecretkey&c=#some,#channels``
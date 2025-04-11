# Basic TCP chat server

This server uses TCP to establish a small insecure chat room with no authentication. I worked on this a bit just to get a feel for different go features.

### Usage

You'll need to have telnet installed. On mac:

```bash
brew install telnet
```

Once installed, start your server:

```bash
go run server.go
```

In as many different terminal windows as you want:

```bash
telnet localhost 8080
```

Each individual terminal window will now be able to communicate with each other in a chat room. Users are notified when new connections are established and when users disconnect.

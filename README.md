# Broadcast Server
A server that can server messages to connected clients. It allows clients to connect to it, send messages that will be broadcasted to all connected clients.

## Features
- It's a command line tool
  - `broadcast start` - starts the server
  - `broadcast stop` - stops the server
  - `broadcast client-ls` - lists all the connected clients
  - `broadcast remove client` - removes a client from receiving messages
- The messages are persisted in a db
  - `broadcast messages-ls` - returns all the historical messages from a client


## Technologies
- Golang
- Websocket
- Go Routines

## Libraries
| Libraries         | Purpose                     |
|-------------------|-----------------------------|
| Cobra             | For CLI                     |
| Gorilla/websocket | For implementing Websockets |



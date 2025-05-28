# Broadcast Server
A server that can server messages to connected clients. It allows clients to connect to it, send messages that will be broadcasted to all connected clients.

## Features
- It's a command line tool
  - `broadcast start` - starts the server
  - `broadcast stop` - stops the server
  - `broadcast version` - version of the server
  - `broadcast clients` - lists all the historical clients
  - `broadcast clients -o` - lists all the online clients
  - `broadcast messages` - returns all the historical messages
  - `broadcast cleandb` - deletes all records of the broadcast server from the database.


## Technologies
- Golang
- Websocket
- SQL
- Go Routines

## Libraries
| Libraries         | Purpose         |
|-------------------|-----------------|
| Cobra             | For CLI         |
| Gorilla/websocket | For Websockets  |      



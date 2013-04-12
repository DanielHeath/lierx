package sockets

import (
  "code.google.com/p/go.net/websocket"
)

type Connection struct {
  ws   *websocket.Conn
  send chan string
}

func (c *connection) reader() {
  for {
    var message string
    err := websocket.Message.Receive(c.ws, &message)
    if err != nil {
      break
    }
    h.broadcast <- message
  }
  c.ws.Close()
}

func (c *connection) writer() {
  for message := range c.send {
    err := websocket.Message.Send(c.ws, message)
    if err != nil {
      break
    }
  }
  c.ws.Close()
}

func wsHandler(room Room) func(*websocket.Conn) {
  return func(*websocket.Conn) {
    c := &connection{send: make(chan string, 256), ws: ws}
    room.register <- c
    defer func() { room.unregister <- c }()
    go c.writer()
    c.reader()
  }
}

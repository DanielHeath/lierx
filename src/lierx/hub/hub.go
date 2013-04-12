// Hub: stolen wholesale from https://gist.gitHub.com/garyburd/1316852
package hub

import (
  "code.google.com/p/go.net/websocket"
)

type Hub struct {
  // Registered Connections.
  Connections map[*Connection]bool

  // Inbound messages from the Connections.
  broadcast chan string

  // Register requests from the Connections.
  register chan *Connection

  // Unregister requests from Connections.
  unregister chan *Connection

  // Send the next frame
  nextFrame func() string
}

func (h *Hub) run() {
  for {
    select {
    case c := <-h.register:
      h.Connections[c] = true
    case c := <-h.unregister:
      delete(h.Connections, c)
      close(c.Send)
    case m := <-h.broadcast:
      for c := range h.Connections {
        select {
        case c.Send <- m:
        default:
          delete(h.Connections, c)
          close(c.Send)
          go c.Ws.Close()
        }
      }
    }
  }
}

type Connection struct {
  Hub
  // The websocket Connection.
  Ws *websocket.Conn

  // Buffered channel of outbound messages.
  Send chan string
}

func (c *Connection) reader() {
  for {
    var message string
    err := websocket.Message.Receive(c.Ws, &message)
    if err != nil {
      break
    }
    c.Hub.broadcast <- c.Hub.nextFrame() // message
  }
  c.Ws.Close()
}

func (c *Connection) writer() {
  for message := range c.Send {
    err := websocket.Message.Send(c.Ws, message)
    if err != nil {
      break
    }
  }
  c.Ws.Close()
}

var H = hub.Hub{
  broadcast:   make(chan string),
  register:    make(chan *connection),
  unregister:  make(chan *connection),
  connections: make(map[*connection]bool),
  nextFrame:   nextFrame(),
}

func wsHandler() func(ws *websocket.Conn) {
  return func(ws *websocket.Conn) {
    c := &Connection{Send: make(chan string, 256), Ws: ws, Hub: H}
    H.register <- c
    defer func() { H.unregister <- c }()
    go c.writer()
    c.reader()
  }
}

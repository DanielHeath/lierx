package sockets

import (
  "code.google.com/p/go.net/websocket"
)

type Room struct {
  // Registered Connections.
  connections map[*Connection]bool

  // Inbound messages from the Connections.
  broadcast chan string

  // Register requests from the Connections.
  register chan *Connection

  // Unregister requests from Connections.
  unregister chan *Connection
}

func (r Room) Send
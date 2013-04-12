package state_sender

type StateFn func() string

func Sender(statefn StateFn) {
  foo
}

type room struct {
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

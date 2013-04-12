package main

import (
  "code.google.com/p/go.net/websocket"
  "fmt"
  "io"
  "lierx/player"
  "log"
  "net/http"
  "strings"
  "time"
)

const (
  gamestate = `{
  "msPerFrame": %d,
  "worms": {%s},
  "map": {
    "baseImg": "/map1.png",
    "width": 1600,
    "height": 1200
  }
}`
)

var players map[string]*player.Player
var msPerFrame = 50

func init() {
  players = make(map[string]*player.Player, 4)
}

func nextFrame() string {
  wormDescriptions := make([]string, 0)
  for _, player := range players {
    player.Tick()
    wormDescriptions = append(wormDescriptions, player.ToJson())
  }
  v := fmt.Sprintf(gamestate, msPerFrame, strings.Join(wormDescriptions, ","))
  return v
}

func ReadString(from io.Reader) string {
  buf := make([]byte, 1024)
  n, err := from.Read(buf)
  if err != nil {
    log.Println(err.Error())
    return ""
  }
  return string(buf[:n])
}

func GameServer(ws *websocket.Conn) {
  name := ws.RemoteAddr().String()
  p := player.NewPlayer(name)
  players[name] = p

  go func() {
    for true {
      str := ReadString(ws)
      if len(str) >= 4 {
        p.ParseControls(str)
      }
    }
  }()

  t := time.NewTicker(time.Millisecond * time.Duration(msPerFrame))
  for _ = range t.C {
    // Each tick, write the current gamestate
    io.WriteString(ws, nextFrame())
  }

}

func main() {
  mux := http.NewServeMux()

  mux.Handle("/ws/", websocket.Handler(GameServer))

  serveStatic := http.FileServer(http.Dir("jekyll/_site"))
  mux.Handle("/", serveStatic)

  log.Fatal(http.ListenAndServe(":8888", mux))
}

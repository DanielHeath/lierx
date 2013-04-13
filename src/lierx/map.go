package main

import (
  _ "image/png"
  "lierx/game"
  "log"
)

func init() {
  game, err := game.NewGame("jekyll/map1.png")
  if err != nil {
    panic(err)
  }
  game.Tick()
  log.Println(game.ToJson())
  panic("HI")
}

---
language: coffeescript
---

@lierx = lierx = {}

lierx.PI = 3.141592653589793
lierx.keyPresses = keyPresses = {}

gameState = {
  worms: {
    "@MichaelDeWildt": {
      x: 89,
      y: 74,
      facingAngle: 149,
    },
    "@Tibbo": {
      x: 1039,
      y: 180,
      facingAngle: 44,
    }
  }
  map: {
    baseImg: "/map1.png"
    width: 1600
    height: 1200
  }
}

$ ->
  $(window).keydown (e) ->
    keyPresses[e.keyCode] = true
  $(window).keyup (e) ->
    delete keyPresses[e.keyCode]

  # Setup paper map
  tx = $("#game")
  width = tx.width()
  height = tx.height()
  tx.replaceWith("<div id='game'></div>")
  tx = $("#game").width(width).height(height)
  window.paper = paper = Raphael(tx[0], tx.width(), tx.height())

  lierx.renderGameState = (state) ->
    maptoPaperX = (x) ->
      (x / state.map.width) * paper.width
    maptoPaperY = (y) ->
      (y / state.map.width) * paper.width

    paper.image("/map1.png", 10, 10, paper.width - 20, paper.height - 20)
    for name, worm of state.worms
      worm.x = maptoPaperX(worm.x)
      worm.y = maptoPaperY(worm.y)
      paper.image("/worm.png", worm.x, worm.y, maptoPaperX(32), maptoPaperY(64))
      paper.text(worm.x, worm.y, name)

  lierx.renderGameState(gameState)

---
language: coffeescript
---

@lierx = lierx = {}

lierx.PI = 3.141592653589793
lierx.keyPresses = keyPresses = {}

$ ->
  $(window).keydown (e) ->
    console.log(e.keyCode)
    keyPresses[e.keyCode] = true
  $(window).keyup (e) ->
    console.log(e.keyCode)
    delete keyPresses[e.keyCode]

  # Setup paper map
  tx = $("#game")
  width = tx.width()
  height = tx.height()
  tx.replaceWith("<div id='game'></div>")
  tx = $("#game").width(width).height(height)
  window.paper = paper = Raphael(tx[0], tx.width(), tx.height())

  lierx.localState = {worms: {}}
  lierx.renderGameState = (state) ->
    maptoPaperX = (x) ->
      (x / state.map.width) * paper.width
    maptoPaperY = (y) ->
      (y / state.map.width) * paper.width

    lierx.localState.map ||= paper.image(state.map.baseImg, 10, 10, paper.width - 20, paper.height - 20)

    for name, worm of state.worms
      worm.x = maptoPaperX(worm.x)
      worm.y = maptoPaperY(worm.y)

      if not lierx.localState.worms[name]
        lierx.localState.worms[name] = {
          img: paper.image("/worm.png", worm.x, worm.y, maptoPaperX(32), maptoPaperY(64))
          text: paper.text(worm.x, worm.y, name)
        }
      else
        lierx.localState.worms[name].img.animate({x: worm.x, y: worm.y}, state.msPerFrame)
        lierx.localState.worms[name].text.animate({x: worm.x, y: worm.y}, state.msPerFrame)



  window.c = c = new WebSocket("ws://localhost:8888/ws/")

  c.onmessage = (me) ->
    lierx.renderGameState(JSON.parse me.data)
    k = (n) -> (if keyPresses[n] then '1' else '0')
    c.send(k(87) + k(83) + k(65) + k(68))


package player

import (
  "fmt"
  "math/rand"
)

const (
  accellerationFromGravity = 9.8 // m/s/s
  accellerationFromRunning = 7   // m/s/s
  friction                 = 2
  jsonFmt                  = `
  "%s": {
    "x": %d,
    "y": %d,
    "facingAngle": %d
  }
  `
)

type Player struct {
  Input
  X     int
  Y     int
  Angle int
  Dx    int
  Dy    int
  Name  string
}

type Input struct {
  Up    bool
  Down  bool
  Left  bool
  Right bool
}

func NewPlayer(name string) *Player {
  return &Player{
    X:     rand.Intn(1600),
    Y:     rand.Intn(1200),
    Angle: rand.Intn(360),
    Name:  name,
    Dx:    rand.Intn(20) - 10,
    Dy:    rand.Intn(20) - 10,
  }
}

func (p Player) ToJson() string {
  return fmt.Sprintf(jsonFmt, p.Name, p.X, p.Y, p.Angle)
}

func (p *Player) Tick() {
  // Add Dx/Dy twice so that degraded framerates don't allow overly high jumping
  p.X += p.Dx
  p.Y += p.Dy
  if p.Input.Left {
    p.Dx -= accellerationFromRunning
  }
  if p.Input.Right {
    p.Dx += accellerationFromRunning
  }

  switch {
  case p.Dx > friction:
    p.Dx -= friction
  case p.Dx < -friction:
    p.Dx += friction
  case true:
    p.Dx = 0
  }

  p.X += p.Dx
  p.Y += p.Dy
}

func (p *Player) ParseControls(form string) {
  p.Input.Up = form[0] == '1'
  p.Input.Down = form[1] == '1'
  p.Input.Left = form[2] == '1'
  p.Input.Right = form[3] == '1'
}

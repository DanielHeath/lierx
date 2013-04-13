package game

import (
  "image"
  "os"
)

type vector struct {
  x     uint
  y     uint
  angle uint8
}

type position struct {
  x uint
  y uint
}
type shape struct {
  position
  width  uint
  height uint
  hitbox [][]byte
}

type entity struct {
  id   uint
  name string
  shape
  mass          uint
  velocity      vector
  accelleration vector
}

type Point struct {
  Indestructable bool
  Passable       bool

  r uint32
  g uint32
  b uint32
  a uint32
}

type Game struct {
  board    [][]Point
  entities []entity
}

func NewGame(mapfile string) (*Game, error) {
  file, err := os.Open(mapfile)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  img, _, err := image.Decode(file)
  if err != nil {
    return nil, err
  }

  XSize := img.Bounds().Dx()
  YSize := img.Bounds().Dy()
  board := make([][]Point, XSize)
  pixels := make([]Point, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
  // Loop over the rows, slicing each row from the front of the remaining pixels slice.
  for i := range board {
    board[i], pixels = pixels[:YSize], pixels[YSize:]
    for j := 0; j < YSize; j++ {
      p := board[i][j]
      p.r, p.g, p.b, p.a = img.At(i, j).RGBA()
      if p.r == 0 && p.g == 0 && p.b == 0 {
        board[i][j].Indestructable = true
      }
      if p.r == 65535 && p.g == 65535 && p.b == 65535 {
        board[i][j].Passable = true
      }
    }
  }

  entities := make([]entity, 0)
  return &Game{board, entities}, nil
}

// Physics:
// 'real-ish' 2d physics isn't too hard.
// for a given def'n
// three-phase:
// 1) is anything intersecting already?
// - panic; exclusion principle. // revisit
// 1a) Is anything moving fast enough to not be adjacent to its previous hitbox?
// - panic; no teleportation
// 2) is anything about to collide (given current velocity)
// - alter its velocity vector to account for the collision
// 3) apply accelleration to velocity

// Integer operations on pixel addresses, but X/Y values can be fractional?
// How about 100 points per pixel.
func (g *Game) Tick() {
}

func (g *Game) ToJson() string {
  return "HI"
}

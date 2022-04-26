package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

// Game implements ebiten.Game interface.
type Game struct{}

var white = color.White

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
    // Write your game's logical update.
    return nil
}

type Point struct {
    x int 
    y int
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
    // Write your game's rendering.
    color := color.White

    points := []Point{{x: 10, y: 12}, {x: 9, y: 12}, {x: 10, y: 11}, {x: 9, y: 11}}

    for _, point := range points {
        screen.Set(point.x, point.y, color)
    }

    line(screen, Point{20, 0}, Point{50, 0})
    line(screen, Point{80, 0}, Point{50, 0})
    line(screen, Point{20, 20}, Point{80, 20})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 320, 240
}

func main() {
    game := &Game{}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Your game's title")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}

type Direction struct {
    right bool 
    up bool
}

func line(screen *ebiten.Image, from, to Point) {
    distanceX := to.x - from.x
    distanceY := to.y - from.y

    direction := Direction{right: distanceX > 0, up: distanceY > 0}

    if distanceX < 0 {
        distanceX *= -1
    }
    if distanceY < 0 {
        distanceY *= -1
    }

    if distanceY == 0 && direction.right {
        for i := distanceX; i > 0; i-- {
            screen.Set(to.x - i, to.y, white)
        }
    } else if distanceY == 0 && !direction.right {
        for i := distanceX; i > 0; i-- {
            screen.Set(to.x + i, to.y, white)
        }
            
    }
}
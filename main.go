package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	if ebiten.IsDrawingSkipped() {
		return
	}
	// Write your game's rendering.
	ebitenutil.DebugPrint(screen, "Hello World")
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 120
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("GG2D")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

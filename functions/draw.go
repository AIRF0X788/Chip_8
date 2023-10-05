package functions

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Chip8) Draw(screen *ebiten.Image) {
	for row := 0; row < gfxHeight; row++ {
		for col := 0; col < gfxWidth; col++ {
			isOn := g.GFX[(row/gfxMultiplier)*chscreenWidth+(col/gfxMultiplier)] == 1
			var colorToUse color.Color
			if isOn {
				colorToUse = color.White
			} else {
				colorToUse = color.Black
			}
			screen.Set(col, row, colorToUse)
		}
	}
}

//

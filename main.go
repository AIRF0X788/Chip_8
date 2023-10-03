package main

import (
	"chip/functions"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const gfxMultiplier = 10
const chscreenWidth = 64
const chscreenHeight = 32
const gfxWidth = chscreenWidth * gfxMultiplier
const gfxHeight = chscreenHeight * gfxMultiplier

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: goc8 [c8 file name]")
	}

	programName := os.Args[1]

	chip := functions.NewChip8()
	chip.LoadApplication(programName)

	ebiten.SetWindowSize(gfxWidth, gfxHeight)
	ebiten.SetWindowTitle(programName)
	if err := ebiten.RunGame(chip); err != nil {
		log.Fatal(err)
	}
}

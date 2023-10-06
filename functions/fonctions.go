package functions

import (
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

func (c *Chip8) updateTimers() {
	if c.DelayTimer > 0 {
		c.DelayTimer--
	}
	if c.SoundTimer > 0 {
		c.SoundTimer--
	}
}

func (g *Chip8) Update() error {
	g.input()
	g.emulateCycle()
	return nil
}

func randomByte() uint8 {
	rand.Seed(time.Now().UTC().UnixNano())
	randint := rand.Intn(math.MaxUint8)
	return uint8(randint)
}

func (c *Chip8) LoadApplication(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file %v", filename)
	}
	copy(c.Memory[512:], bytes)
}

func (g *Chip8) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gfxWidth, gfxHeight
}

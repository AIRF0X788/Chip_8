package data

const chscreenWidth = 64
const chscreenHeight = 32

type Chip8 struct {
	Opcode     uint16
	Memory     [4096]uint8
	V          [16]uint8
	I          uint16
	PC         uint16
	GFX        [chscreenWidth * chscreenHeight]uint8
	DelayTimer uint8
	SoundTimer uint8
	Stack      [16]uint16
	SP         uint16
	Key        [16]bool
}

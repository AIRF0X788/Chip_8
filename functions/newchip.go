package functions

func NewChip8() *Chip8 {
	chip := &Chip8{
		PC: 0x200,
	}
	copy(chip.Memory[:], fontSet[:])
	return chip
}

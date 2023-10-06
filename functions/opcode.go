package functions

import "log"

func (c *Chip8) fetchOpcode() {
	c.Opcode = uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1])
}

func (c *Chip8) executeOpcode() bool {
	switch c.Opcode & 0xF000 {
	case 0x0000:
		switch c.Opcode & 0x000F {
		case 0x0000:
			c.GFX = [2048]uint8{}
			c.PC += 2
		case 0x000E:
			c.SP--
			c.PC = c.Stack[c.SP]
			c.PC += 2
		default:
			panicUnknownOpcode(c.Opcode)
		}
	case 0x1000:
		c.PC = c.Opcode & 0x0FFF
	case 0x2000:
		c.Stack[c.SP] = c.PC
		c.SP++
		c.PC = c.Opcode & 0x0FFF
	case 0x3000:
		if c.V[(c.Opcode&0x0F00)>>8] == (uint8(c.Opcode) & 0x00FF) {
			c.PC += 4
		} else {
			c.PC += 2
		}
	case 0x4000:
		if c.V[(c.Opcode&0x0F00)>>8] != (uint8(c.Opcode) & 0x00FF) {
			c.PC += 4
		} else {
			c.PC += 2
		}
	case 0x5000:
		if c.V[(c.Opcode&0x0F00)>>8] != c.V[(uint8(c.Opcode)&0x00F0)>>4] {
			c.PC += 4
		} else {
			c.PC += 2
		}
	case 0x6000:
		c.V[(c.Opcode&0x0F00)>>8] = uint8(c.Opcode) & 0x00FF
		c.PC += 2
	case 0x7000:
		c.V[(c.Opcode&0x0F00)>>8] += uint8(c.Opcode) & 0x00FF
		c.PC += 2
	case 0x8000:
		switch c.Opcode & 0x000F {
		case 0x0000:
			c.V[(c.Opcode&0x0F00)>>8] = c.V[(c.Opcode&0x00F0)>>4]
			c.PC += 2
		case 0x0001:
			c.V[(c.Opcode&0x0F00)>>8] |= c.V[(c.Opcode&0x00F0)>>4]
			c.PC += 2
		case 0x0002:
			c.V[(c.Opcode&0x0F00)>>8] &= c.V[(c.Opcode&0x00F0)>>4]
			c.PC += 2
		case 0x0003:
			c.V[(c.Opcode&0x0F00)>>8] ^= c.V[(c.Opcode&0x00F0)>>4]
			c.PC += 2
		case 0x0004:
			if c.V[(c.Opcode&0x00F0)>>4] > (0xFF - c.V[(c.Opcode&0x0F00)>>8]) {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.V[(c.Opcode&0x0F00)>>8] += c.V[(c.Opcode&0x00F0)>>4]
			c.PC += 2
		case 0x0005:
			if c.V[(c.Opcode&0x00F0)>>4] > c.V[(c.Opcode&0x0F00)>>8] {

				c.V[0xF] = 0
			} else {

				c.V[0xF] = 1
			}
			c.V[(c.Opcode&0x0F00)>>8] -= c.V[(c.Opcode&0x00F0)>>4]
			c.PC += 2

		case 0x0006:
			c.V[0xF] = c.V[(c.Opcode&0x0F00)>>8] & 0x1
			c.V[(c.Opcode&0x0F00)>>8] >>= 1
			c.PC += 2
		case 0x0007:
			if c.V[(c.Opcode&0x0F00)>>8] > c.V[(c.Opcode&0x00F0)>>4] {
				c.V[0xF] = 0
			} else {
				c.V[0xF] = 1
			}
			c.V[(c.Opcode&0x0F00)>>8] = c.V[(c.Opcode&0x00F0)>>4] - c.V[(c.Opcode&0x0F00)>>8]
			c.PC += 2
		case 0x000E:
			c.V[0xF] = c.V[(c.Opcode&0x0F00)>>8] >> 7
			c.V[(c.Opcode&0x0F00)>>8] <<= 1
			c.PC += 2
		default:
			panicUnknownOpcode(c.Opcode)
		}
	case 0x9000:
		if c.V[(c.Opcode&0x0F00)>>8] != c.V[(c.Opcode&0x00F0)>>4] {
			c.PC += 4
		} else {
			c.PC += 2
		}
	case 0xA000:
		c.I = c.Opcode & 0x0FFF
		c.PC += 2
	case 0xB000:
		c.PC = (c.Opcode & 0x0FFF) + uint16(c.V[0])
	case 0xC000:
		c.V[(c.Opcode&0x0F00)>>8] = randomByte() & uint8(c.Opcode&0x00FF)
		c.PC += 2
	case 0xD000:
		x := uint16(c.V[(c.Opcode&0x0F00)>>8])
		y := uint16(c.V[(c.Opcode&0x00F0)>>4])
		height := uint16(c.Opcode & 0x000F)
		var pixel uint16

		c.V[0xF] = 0

		for yline := uint16(0); yline < height; yline++ {
			pixel = uint16(c.Memory[c.I+yline])
			for xline := uint16(0); xline < 8; xline++ {
				if (pixel & (0x80 >> xline)) != 0 {
					if c.GFX[x+xline+((y+yline)*chscreenWidth)] == 1 {
						c.V[0xF] = 1
					}
					c.GFX[x+xline+((y+yline)*chscreenWidth)] ^= 1
				}
			}
		}

		c.PC += 2

	case 0xE000:
		switch c.Opcode & 0x00FF {
		case 0x009E:
			if c.Key[c.V[(c.Opcode&0x0F00)>>8]] {
				c.PC += 4
			} else {
				c.PC += 2
			}
		case 0x00A1:
			if !c.Key[c.V[(c.Opcode&0x0F00)>>8]] {
				c.PC += 4
			} else {
				c.PC += 2
			}
		default:
			panicUnknownOpcode(c.Opcode)
		}
	case 0xF000:
		switch c.Opcode & 0x00FF {
		case 0x0007:
			c.V[(c.Opcode&0x0F00)>>8] = c.DelayTimer
			c.PC += 2
		case 0x000A:
			keyPress := false
			for i := uint8(0); i < 16; i++ {
				if c.Key[i] {
					c.V[(c.Opcode&0x0F00)>>8] = i
					keyPress = true
				}
			}
			if !keyPress {
				return true
			}
			c.PC += 2
		case 0x0015:
			c.DelayTimer = c.V[(c.Opcode&0x0F00)>>8]
			c.PC += 2
		case 0x0018:

			c.SoundTimer = c.V[(c.Opcode&0x0F00)>>8]
			c.PC += 2
		case 0x001E:
			if c.I+uint16(c.V[(c.Opcode&0x0F00)>>8]) > 0xFFF {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.I += uint16(c.V[(c.Opcode&0x0F00)>>8])
			c.PC += 2
		case 0x0029:
			c.I = uint16(c.V[(c.Opcode&0x0F00)>>8]) * 0x5
			c.PC += 2
		case 0x0033:
			c.Memory[c.I] = c.V[(c.Opcode&0x0F00)>>8] / 100
			c.Memory[c.I+1] = (c.V[(c.Opcode&0x0F00)>>8] / 10) % 10
			c.Memory[c.I+2] = (c.V[(c.Opcode&0x0F00)>>8] % 100) % 10
			c.PC += 2
		case 0x0055:
			for i := uint16(0); i <= ((c.Opcode & 0x0F00) >> 8); i++ {
				c.Memory[c.I+i] = c.V[i]
			}
			c.I += ((c.Opcode & 0x0F00) >> 8) + 1
			c.PC += 2
		case 0x0065:
			for i := uint16(0); i <= ((c.Opcode & 0x0F00) >> 8); i++ {
				c.V[i] = c.Memory[c.I+i]
			}
			c.I += ((c.Opcode & 0x0F00) >> 8) + 1
			c.PC += 2

		default:
			panicUnknownOpcode(c.Opcode)
		}
	default:
		panicUnknownOpcode(c.Opcode)
	}
	return false

}

func panicUnknownOpcode(opcode uint16) {
	log.Panicf("Unknown opcode %v", opcode)
}

package functions

func (c *Chip8) emulateCycle() {
	c.fetchOpcode()
	skip := c.executeOpcode()
	if skip {
		return
	}
	c.updateTimers()
}

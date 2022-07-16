package command

// Concrete Command

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

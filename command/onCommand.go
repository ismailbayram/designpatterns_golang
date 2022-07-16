package command

// Concrete Command

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

package command

// Receiver Interface

type device interface {
	on()
	off()
}

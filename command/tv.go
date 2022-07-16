package command

import "fmt"

// Concrete Receiver

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("TV is ON")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("TV is OFF")
}

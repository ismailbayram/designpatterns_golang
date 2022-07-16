package adapter

import "fmt"

type mac struct {
}

func (m *mac) insertIntoLightingPort() {
	fmt.Println("Lighting connector is plugged into MAC machine")
}

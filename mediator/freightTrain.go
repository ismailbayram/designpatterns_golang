package mediator

import "fmt"

// Concrete Component

type freightTrain struct {
	mediator mediator
}

func (g *freightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("Freight Train: ARRIVAL BLOCKED!")
		return
	}
	fmt.Println("FreightTrain: ARRIVED")
}

func (g *freightTrain) depart() {
	fmt.Println("FreightTrain: LEAVING")
	g.mediator.notifyAboutDeparture()
}

func (g *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, ARRIVING")
	g.arrive()
}

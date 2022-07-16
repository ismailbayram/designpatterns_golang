package mediator

import "fmt"

// Concrete Component

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("Passenger Train: ARRIVAL BLOCKED!")
		return
	}
	fmt.Println("PassengerTrain: ARRIVED")
}

func (g *passengerTrain) depart() {
	fmt.Println("PassengerTrain: LEAVING")
	g.mediator.notifyAboutDeparture()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, ARRIVING")
	g.arrive()
}

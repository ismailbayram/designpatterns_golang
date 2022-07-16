package mediator

func RunMediator() {
	stationManager := newStationManager()

	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}

	freightTrain := &freightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}

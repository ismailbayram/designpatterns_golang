package mediator

// Mediator Interface

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

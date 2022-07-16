package mediator

// Component

type train interface {
	arrive()
	depart()
	permitArrival()
}

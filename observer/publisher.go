package observer

type publisher interface {
	subcribe(Observer observer)
	unsubscribe(Observer observer)
	notifyAll()
}

package observer

// Observer Interface

type observer interface {
	update(string)
	getID() string
}

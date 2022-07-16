package observer

func RunObserver() {
	shirtItem := newItem("Shirt")

	observer1 := &customer{id: "observer 1"}
	observer2 := &customer{id: "observer 2"}

	shirtItem.subscribe(observer1)
	shirtItem.subscribe(observer2)

	shirtItem.updateAvailability()

	shirtItem.unsubscribe(observer2)
	shirtItem.updateAvailability()
}

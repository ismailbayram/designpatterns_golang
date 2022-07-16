package decorator

import "fmt"

func RunDecorator() {
	pizza := &veggeMania{}

	// cheese topping wraps the vegge mania
	pizzaWithChees := &cheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithChees,
	}

	fmt.Printf("price of all mixup: %d\n", pizzaWithCheeseAndTomato.getPrice())
}

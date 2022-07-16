package observer

import "fmt"

// Concrete Observer

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}

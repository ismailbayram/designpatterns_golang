package decorator

type veggeMania struct {
}

// concrete component

func (p *veggeMania) getPrice() int {
	return 15
}

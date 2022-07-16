package factory

// Product Interface

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

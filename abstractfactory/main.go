package abstractfactory

import "fmt"

func RunAbstractFactory() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	fmt.Println(adidasShoe.getSize(), adidasShoe.getLogo())
	fmt.Println(adidasShirt.getSize(), adidasShirt.getLogo())
	fmt.Println(nikeShoe.getSize(), nikeShoe.getLogo())
	fmt.Println(nikeShirt.getSize(), nikeShirt.getLogo())
}

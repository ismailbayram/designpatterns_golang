package factory

import "fmt"

func RunFactory() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	fmt.Println(ak47.getName(), ak47.getPower())
	fmt.Println(musket.getName(), musket.getPower())
}

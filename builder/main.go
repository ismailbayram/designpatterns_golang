package builder

import "fmt"

func RunBuilder() {
	normalBuilder, _ := getBuilder("normal")
	iglooBuilder, _ := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Println(normalHouse.windowType, normalHouse.floor, normalHouse.doorType)

	director.setBuilder(iglooBuilder)
	iglooHosue := director.buildHouse()
	fmt.Println(iglooHosue.windowType, iglooHosue.floor, iglooHosue.doorType)

}

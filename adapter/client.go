package adapter

import "fmt"

type client struct {
}

func (c *client) insertLightingConnectorIntocomputer(com computer) {
	fmt.Println("client inserts lighting connector into computer")
	com.insertIntoLightingPort()
}

package adapter

func RunAdapter() {
	client := &client{}
	macMachine := &mac{}

	client.insertLightingConnectorIntocomputer(macMachine)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}
	client.insertLightingConnectorIntocomputer(windowsMachineAdapter)
}

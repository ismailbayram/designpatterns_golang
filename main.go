package main

import (
	"design_patterns/abstractfactory"
	"design_patterns/adapter"
	"design_patterns/builder"
	"design_patterns/command"
	"design_patterns/composite"
	"design_patterns/decorator"
	"design_patterns/factory"
	"design_patterns/mediator"
	"design_patterns/observer"
	"design_patterns/state"
	"fmt"
)

func main() {
	abstractfactory.RunAbstractFactory()
	fmt.Println("********************************************")
	factory.RunFactory()
	fmt.Println("********************************************")
	builder.RunBuilder()
	fmt.Println("********************************************")
	adapter.RunAdapter()
	fmt.Println("********************************************")
	composite.RunComposite()
	fmt.Println("********************************************")
	decorator.RunDecorator()
	fmt.Println("********************************************")
	command.RunCommand()
	fmt.Println("********************************************")
	mediator.RunMediator()
	fmt.Println("********************************************")
	observer.RunObserver()
	fmt.Println("********************************************")
	state.RunState()
	fmt.Println("********************************************")
}

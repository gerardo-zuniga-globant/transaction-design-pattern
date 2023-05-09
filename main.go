package main

import (
	"fmt"
	"time"
)

type Flow interface {
	validate(Input) error
	process(Input, CommonType) error
}

// There are 9 Types(Services) to attend, that provoke multiple SubTypes per each validation service execution.
func (i Input) validateType() {
	switch i.MyType {
	case Service_Deposit:
		ServiceDeposit{}.validate(i)
	case Service_TicketZ:
		ServiceTicketZ{}.validate(i)
	}
}

func (i Input) processByType() {
	switch i.MyType {
	case Service_Deposit:

	case Service_TicketZ:

	}
}

func (s ServiceDeposit) validate(input Input) error {
	fmt.Println("I'm validating Deposit and doing something special ")
	ct := CommonType{
		CommonAtt: "I'm common from Deposit Service",
	}
	s2ST1 := ServiceDeposit_SubType1{
		Att1: "I'm subType1 from Deposit",
	}
	t2ST2 := ServiceDeposit_SubType2{
		Att1: "I'm subType2 from Deposit",
	}
	fmt.Println(ct, s2ST1, t2ST2)

	return nil
}

func (s ServiceDeposit) process(d Deposit_Data) error {
	fmt.Println("This is the process ServiceDeposit")
	return nil
}

func (s ServiceTicketZ) validate(input Input) error {
	fmt.Println("I'm validating TicketZ type and sleeping zzzzzz ")
	ct := CommonType{
		CommonAtt: "I'm common from Ticket Z",
	}
	t3ST1 := ServiceTicketZ_SubType1{
		Att1: "I'm subType1 from Ticket Z",
	}
	fmt.Println(ct, t3ST1)
	return nil
}

func (s ServiceTicketZ) process(d TicketZ_Data) error {
	fmt.Println("This is the process ServiceTicketZ")
	return nil
}

func main() {
	/*
		Input{
			Att1:   "Hello World",
			MyType: Service_Deposit,
		}.validateType()
		fmt.Println("************************************************************************************")
		Input{
			Att1:   "Hola Mundo",
			MyType: Service_TicketZ,
		}.validateType()*/

	//create channels
	chTicketz := make(chan TicketZ_Data, 50)
	chDeposit := make(chan Deposit_Data, 50)

	//launch channel listeners
	go ServiceTicketZ{}.listenerProcessServiceTicketZ(chTicketz)
	go ServiceDeposit{}.listenerProcessServiceDeposit(chDeposit)

	//an input enter
	example1 := Input{MyType: Service_Deposit}
	Validate(example1, chTicketz, chDeposit)

	example2 := Input{MyType: Service_TicketZ}
	Validate(example2, chTicketz, chDeposit)

	example3 := Input{MyType: Service_Deposit}
	Validate(example3, chTicketz, chDeposit)

	time.Sleep(10 * time.Second)
}

func (s ServiceTicketZ) listenerProcessServiceTicketZ(c <-chan TicketZ_Data) {
	for input := range c {
		s.process(input)
	}
}

func (s ServiceDeposit) listenerProcessServiceDeposit(c <-chan Deposit_Data) {
	for input := range c {
		s.process(input)
	}
}

func Validate(input Input, cTicketz chan<- TicketZ_Data, cDeposit chan<- Deposit_Data) {
	switch input.MyType {
	case Service_TicketZ:
		time.Sleep(2 * time.Second)
		cTicketz <- TicketZ_Data{}
	case Service_Deposit:
		time.Sleep(2 * time.Second)
		cDeposit <- Deposit_Data{}
	}
}

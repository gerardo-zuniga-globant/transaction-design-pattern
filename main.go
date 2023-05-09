package main

import "fmt"

type Flow interface {
	validate(Input) error
	process(Input, CommonType) error
}

func (i Input) validateType() {
	switch i.MyType {
	case Service_1:
		ServiceDeposit{
			RedisService: i.Att1,
		}.validate(i)
	case Service_3:
		ServiceTicketZ{
			DataBaseService: i.Att1,
		}.validate(i)
	}
}

func (i Input) processByType() {
	switch i.MyType {
	case Service_1:

	case Service_3:

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

func (s ServiceDeposit) process(i Input, c CommonType) error {

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

func (s ServiceTicketZ) process(i Input, c CommonType) error {

	return nil
}

func main() {
	Input{
		Att1:   "Hello World",
		MyType: Service_1,
	}.validateType()
	fmt.Println("************************************************************************************")
	Input{
		Att1:   "Hola Mundo",
		MyType: Service_3,
	}.validateType()
}

package main

import "fmt"

type Flow interface {
	validate(Input) error
	process(Input, CommonType) error
}

func (i Input) validateType() {
	switch i.MyType {
	case Service_2:
		Service2{
			Att1: i.Att1,
		}.validate(i)
	case Service_3:
		Service3{
			Att1: i.Att1,
		}.validate(i)
	}
}

func (i Input) processByType() {
	switch i.MyType {
	case Service_2:

	case Service_3:

	}
}

func (s Service2) validate(input Input) error {
	fmt.Println("I'm validating Deposit and doing something special ")
	ct := CommonType{
		CommonAtt: "I'm common from 2nd Type",
	}
	s2ST1 := Type2SubType1{
		Att1: "I'm subType1 from Deposit",
	}
	t2ST2 := Type2SubType2{
		Att1: "I'm subType2 from Deposit",
	}
	fmt.Println(ct, s2ST1, t2ST2)

	return nil
}

func (s Service2) process(i Input, c CommonType) error {

	return nil
}

func (s Service3) validate(input Input) error {
	fmt.Println("I'm validating TicketZ type and sleeping zzzzzz ")
	ct := CommonType{
		CommonAtt: "I'm common from Ticket Z",
	}
	t3ST1 := Type3SubType1{
		Att1: "I'm subType1 from Ticket Z",
	}
	fmt.Println(ct, t3ST1)
	return nil
}

func (s Service3) process(i Input, c CommonType) error {

	return nil
}

func main() {
	Input{
		Att1:   "Hello World",
		MyType: Service_2,
	}.validateType()
	fmt.Println("************************************************************************************")
	Input{
		Att1:   "Hola Mundo",
		MyType: Service_3,
	}.validateType()
}

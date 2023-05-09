package main

import "fmt"

type Transaction interface {
	validate(Input) (interface{}, error) //this is not accomplishing the interface and the testing will be a pain in the ....
	process(Input) error
}

// There are 9 Types(Services) to attend, that provoke multiple SubTypes per each validation Service execution.
func (i Input) validateType() {
	switch i.MyType {
	case Service_Deposit:
		s, _ := ServiceDeposit{
			RedisService:    "redis",
			DataBaseService: "skillz",
		}.validate(i)
		s.process(i)
	case Service_TicketZ:
		s, _ := ServiceTicketZ{
			DataBaseService: "ums",
			Log:             "log",
		}.validate(i)
		s.process(i)
	}
}

func (s ServiceDeposit) validate(input Input) (DepositResult, error) {
	fmt.Println("I'm validating Deposit and doing something special ")
	ct := CommonType{
		CommonAtt: "I'm ValidateCommon from Deposit Service",
	}
	s2ST1 := ServiceDeposit_SubType1{
		Att1: "I'm subType1 from Deposit",
	}
	t2ST2 := ServiceDeposit_SubType2{
		Att1: "I'm subType2 from Deposit",
	}
	fmt.Println(ct, s2ST1, t2ST2)

	return DepositResult{
		Service:        s,
		ValidateCommon: ct,
		ValidateType1:  s2ST1,
		ValidateType2:  t2ST2,
	}, nil
}

func (s DepositResult) process(i Input) error {
	fmt.Println("doing something with", s.Service, " Service from type", i.MyType, " with the values", s.ValidateType1, " and ", s.ValidateType2, " and ValidateCommon value ", s.ValidateCommon)
	return nil
}

func (s ServiceTicketZ) validate(input Input) (TicketZResult, error) {
	fmt.Println("I'm validating TicketZ type and sleeping zzzzzz ")
	ct := CommonType{
		CommonAtt: "I'm ValidateCommon from Ticket Z",
	}
	t3ST1 := ServiceTicketZ_SubType1{
		Att1: "I'm subType1 from Ticket Z",
	}
	fmt.Println(ct, t3ST1)
	return TicketZResult{
		Service: s,
		common:  ct,
		type1:   t3ST1,
	}, nil
}

func (s TicketZResult) process(i Input) error {
	fmt.Println("doing something with", s.Service, " Service from type", i.MyType, " with the values", s.type1, " and ValidateCommon value ", s.common)
	return nil
}

func main() {
	Input{
		Att1:   "Hello World",
		MyType: Service_Deposit,
	}.validateType()
	fmt.Println("************************************************************************************")
	Input{
		Att1:   "Hola Mundo",
		MyType: Service_TicketZ,
	}.validateType()
}

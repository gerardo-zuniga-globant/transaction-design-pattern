package main

// Original Type
const (
	Service_Deposit MyOwnType = "Deposit"
	Service_TicketZ MyOwnType = "Ticket-z"
)

type MyOwnType string

type Input struct {
	Att1   string
	Att2   string
	MyType MyOwnType
}

type Result struct {
	DataProcesed string
}

type SomeStruct struct {
	Data string
}

type CommonType struct {
	CommonAtt string
}

// ServiceDeposit ****************  2nd Type Info
type ServiceDeposit struct {
	RedisService    string
	DataBaseService string
}

type ServiceDeposit_SubType1 struct {
	Att1 string
}

type ServiceDeposit_SubType2 struct {
	Att1 string
}

// ServiceTicketZ ****************  3rd Type Info
type ServiceTicketZ struct {
	DataBaseService string
	Log             string
}

type ServiceTicketZ_SubType1 struct {
	Att1 string
}

package main

// Original Type
const (
	Service_2 MyOwnType = "Deposit"
	Service_3 MyOwnType = "Ticket-z"
)

type MyOwnType string

type Input struct {
	Att1   string
	MyType MyOwnType
}

type CommonType struct {
	CommonAtt string
}

// Service2 ****************  2nd Type Info
type Service2 struct {
	Att1 string
	Att2 bool
}

type Type2SubType1 struct {
	Att1 string
}

type Type2SubType2 struct {
	Att1 string
}

// Service3 ****************  3rd Type Info
type Service3 struct {
	Att1 string
	Att2 bool
}

type Type3SubType1 struct {
	Att1 string
}

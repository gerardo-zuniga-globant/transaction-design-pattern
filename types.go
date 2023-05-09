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

// There are 9 Types(transactions) to attend, that provoke multiple SubTypes per each validation service execution.
func (i Input) createTransactionByType(service *TransactionService) Transaction {
	switch i.MyType {
	case Service_Deposit:
		return &DepositTransaction{
			Service: service,
			Att1:    i.Att1,
		}
	case Service_TicketZ:
		return &TicketZTransaction{
			Service: service,
			Att2:    i.Att2,
		}
	}
	return nil
}

type Result struct {
	DataProcesed string
}

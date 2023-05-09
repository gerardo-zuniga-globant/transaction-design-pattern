package main

import (
	"fmt"
)

func main() {

	repo := &Repo{}
	logger := &Logger{}
	service := NewTransactionService(repo, logger)

	//This will be the request received in the controller
	input := Input{
		Att1:   "Fer",
		MyType: Service_Deposit,
	}

	/*
		input2 := Input{
			Att2:   "Gerard",
			MyType: Service_TicketZ,
		}
	*/

	transaction := input.createTransactionByType(service)

	err := transaction.validate()
	if err != nil {
		fmt.Println("Hubo un error!")
	}

	result, err := transaction.process()
	fmt.Println("Result: ", result.DataProcesed)
}

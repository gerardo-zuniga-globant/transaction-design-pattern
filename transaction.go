package main

import "errors"

type Transaction interface {
	validate() error
	process() (*Result, error)
}

type DepositTransaction struct {
	Service    *TransactionService
	Att1       string
	DataStored string
}

type TicketZTransaction struct {
	Service    *TransactionService
	Att2       string
	DataStored string
}

func (dt *DepositTransaction) validate() error {
	//TODO: Validate logic
	if dt.Att1 == "" {
		return errors.New("asd")
	}
	dt.DataStored = dt.Service.Repo.GetData(dt.Att1)
	return nil
}

func (dt *DepositTransaction) process() (*Result, error) {
	return &Result{
		DataProcesed: dt.DataStored + " - Processed by Deposit TransactionService",
	}, nil
}

func (dt *TicketZTransaction) validate() error {
	//TODO: Validate logic
	if dt.Att2 == "" {
		return errors.New("asd")
	}
	dt.DataStored = dt.Service.Repo.GetData(dt.Att2)
	return nil
}

func (dt *TicketZTransaction) process() (*Result, error) {
	return &Result{
		DataProcesed: dt.DataStored + " - Processed by Ticketz TransactionService",
	}, nil
}

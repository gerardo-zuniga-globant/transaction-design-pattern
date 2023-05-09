package main

import (
	"errors"
	"fmt"
)

type TransactionService struct {
	Repo   *Repo
	Logger *Logger
}

func NewTransactionService(repo *Repo, logger *Logger) *TransactionService {
	return &TransactionService{
		Repo:   repo,
		Logger: logger,
	}
}

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

type Repo struct {
}

func (r *Repo) GetData(param string) string {
	return fmt.Sprintf("Hola %s", param)
}

type Logger struct {
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

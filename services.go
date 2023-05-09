package main

import (
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

type Repo struct {
}

func (r *Repo) GetData(param string) string {
	return fmt.Sprintf("Hola %s", param)
}

type Logger struct {
}

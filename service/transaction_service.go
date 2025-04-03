package service

import "tracker/repository"

type TransactionService struct {
	Repo repository.TransactionRepository
}
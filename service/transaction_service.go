package service

import "tracker/repository"

type TransactionService struct {
	repo repository.TransactionRepository
}
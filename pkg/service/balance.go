package service

import (
	balance_api "balance-api"
	"balance-api/pkg/repository"
	"errors"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetBalance(id int) (int, error) {
	return s.repo.GetBalance(id)
}

func (s *BalanceService) UpdateBalance(id int, input balance_api.UpdateBalance, sourceOrPurpose string) error {

	curBalance, err := s.repo.GetBalance(id)
	if err != nil && input.CashFlow == true {
		return s.repo.CreateBalance(id, input)
	} else if err != nil && input.CashFlow == false {
		return errors.New("You need send some money to this id first")
	} else if curBalance < input.Netto && input.CashFlow == false {
		return errors.New("Balance has not enough money")
	} else {
		return s.repo.UpdateBalance(id, input, sourceOrPurpose)
	}
}

package services

import (
	"kasir-api/dto"
	"kasir-api/models"
	"kasir-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) SalesSummaryToday() (models.RecapToday, error) {
	return s.repo.SalesSummaryToday()
}

func (s *TransactionService) GetTransactionToday() ([]models.Transaction, error) {
	return s.repo.GetTransactionToday()
}
func (s *TransactionService) GetReport(startDate string, endDate string) (dto.ReportResponse, error) {
	return s.repo.GetReport(startDate, endDate)
}

package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(transac models.TransactionModels) (models.TransactionModels, error)
}

func RepositoryTransac(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transac models.TransactionModels) (models.TransactionModels, error) {
	err := r.db.Preload("Trip.Country").Preload("User").Create(&transac).Error

	return transac, err
}

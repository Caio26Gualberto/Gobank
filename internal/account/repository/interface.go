package repository

import (
	"github.com/Caio26Gualberto/gobank/internal/account/models"
)

type AccountRepository interface {
	Create(account *models.Account) (int64, error)
	GetById(id int64) (*models.Account, error)
	List() ([]*models.Account, error)
	Update(account *models.Account) error
	Delete(id int64) error
}

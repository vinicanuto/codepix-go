package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/vinicanuto/codepix/domain/model"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

// type TransactionRepositoryInterface interface {
// 	Register(transaction *Transaction) error
// 	Save(transaction *Transaction) error
// 	Find(id string) (*Transaction, error)
// }

func (t TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error
	return err
}

func (t TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error
	return err
}

func (t TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {

	var transaction model.Transaction

	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}

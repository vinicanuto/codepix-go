package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status"  valid:"notnull"`
}

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
	FindBank(id string) (*Bank, error)
}

func (p *PixKey) isValid() error {

	_, err := govalidator.ValidateStruct(p)

	if p.Kind != "email" && p.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if p.Status != "active" && p.Status != "inactive" {
		return errors.New("invalid status")
	}

	return err
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:      kind,
		Account:   account,
		AccountID: account.ID,
		Key:       key,
		Status:    "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}

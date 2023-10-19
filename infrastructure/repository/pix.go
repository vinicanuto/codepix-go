package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/vinicanuto/codepix/domain/model"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (p PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := p.Db.Create(bank).Error
	return err
}

func (p PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := p.Db.Create(account).Error
	return err
}

func (p PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := p.Db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (p PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	pixKey.ID = key
	pixKey.Kind = kind

	p.Db.Preload("Account.Bank").First(&pixKey, "key = ? and kind = ?", key, kind)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

func (p PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	p.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}

	return &account, nil
}

func (p PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	p.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}

	return &bank, nil
}

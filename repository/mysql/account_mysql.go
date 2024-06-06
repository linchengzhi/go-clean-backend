package mysql

import (
	"context"

	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return AccountRepo{db: db}
}

func (rp *AccountRepo) WithTx(tx *gorm.DB) *AccountRepo {
	return &AccountRepo{db: tx}
}

func (rp *AccountRepo) GetByID(ctx context.Context, id int) (*entity.Account, error) {
	var account *entity.Account
	err := rp.db.First(&account, id).Error
	return account, errors.Wrap(err, "failed to get account by id")
}

func (rp *AccountRepo) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account *entity.Account
	err := rp.db.Where("email = ?", email).First(&account).Error
	return account, errors.Wrap(err, "failed to get account by email")
}

func (rp *AccountRepo) Create(ctx context.Context, account *entity.Account) error {
	err := rp.db.Create(account).Error
	return errors.Wrap(err, "failed to create account")
}

func (rp *AccountRepo) Update(ctx context.Context, account *entity.Account) error {
	err := rp.db.Save(account).Error
	return errors.Wrap(err, "failed to update account")
}

func (rp *AccountRepo) Delete(ctx context.Context, id int) error {
	err := rp.db.Delete(&entity.Account{}, id).Error
	return errors.Wrap(err, "failed to delete account")
}

package mysql

import (
	"context"

	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db: db}
}

func (rp *UserRepo) WithTx(tx *gorm.DB) *UserRepo {
	return &UserRepo{db: tx}
}

func (rp *UserRepo) GetById(ctx context.Context, id int) (*entity.User, error) {
	var user *entity.User
	err := rp.db.First(&user, id).Error
	return user, err
}

func (rp *UserRepo) GetByAccountId(ctx context.Context, id int) (*entity.User, error) {
	var user *entity.User
	err := rp.db.Where("account_id = ?", id).First(&user).Error
	return user, err
}

func (rp *UserRepo) Create(ctx context.Context, user *entity.User) error {
	err := rp.db.Create(user).Error
	return err
}

func (rp *UserRepo) Save(ctx context.Context, user *entity.User) error {
	err := rp.db.Save(user).Error
	return err
}

func (rp *UserRepo) Delete(ctx context.Context, id int) error {
	err := rp.db.Delete(&entity.User{}, id).Error
	return err
}

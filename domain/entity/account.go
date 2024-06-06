package entity

import (
	"context"
	"time"
)

type Account struct {
	Id         int       `json:"id"`
	Password   string    `json:"password"`
	Salt       string    `json:"salt"`
	Email      string    `json:"email"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (a *Account) TableName() string {
	return "account"
}

type IAccountRepo interface {
	GetByID(ctx context.Context, id int) (*Account, error)
	GetByEmail(ctx context.Context, email string) (*Account, error)
	Create(ctx context.Context, account *Account) error
	Update(ctx context.Context, account *Account) error
	Delete(ctx context.Context, id int) error
}

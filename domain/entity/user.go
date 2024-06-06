package entity

import (
	"context"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	AccountId  int       `json:"account_id"`
	Name       string    `json:"name"`
	Birthday   time.Time `json:"birthday"`
	Gender     int       `json:"gender"`
	Profile    string    `json:"profile"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}

type IUserRepo interface {
	GetById(ctx context.Context, id int) (*User, error)
	GetByAccountId(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, user *User) error
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}

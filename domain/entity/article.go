package entity

import (
	"context"
	"time"
)

type Article struct {
	ID         int       `json:"id"`
	Title      string    `json:"title" validate:"required"`
	Content    string    `json:"content" validate:"required"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (u *Article) TableName() string {
	return "article"
}

type IArticleRepo interface {
	GetById(ctx context.Context, id int) (*Article, error)
	ListByTitle(ctx context.Context, title string, page, pageSize int) ([]*Article, error)
	Create(ctx context.Context, user *Article) error
	Update(ctx context.Context, user *Article) error
	Delete(ctx context.Context, id int) error
}

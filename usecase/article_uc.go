package usecase

import (
	"context"
	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"github.com/linchengzhi/go-clean-backend/repository/mysql"
	"time"
)

type IArticleUc interface {
	Add(ctx context.Context, title, content string) error
	GetByID(ctx context.Context, id int) (*entity.Article, error)
	List(ctx context.Context, title string, page, pageSize int) ([]*entity.Article, error)
}

type ArticleUc struct {
	articleRepo mysql.ArticleRepo
}

func NewArticleUc(articleRepo mysql.ArticleRepo) IArticleUc {
	return ArticleUc{
		articleRepo: articleRepo,
	}
}

func (uc ArticleUc) Add(ctx context.Context, title, content string) error {
	article := &entity.Article{
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return uc.articleRepo.Create(ctx, article)
}

func (uc ArticleUc) GetByID(ctx context.Context, id int) (*entity.Article, error) {
	return uc.articleRepo.GetById(ctx, id)
}

func (uc ArticleUc) List(ctx context.Context, title string, page, pageSize int) ([]*entity.Article, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return uc.articleRepo.ListByTitle(ctx, title, page, pageSize)
}

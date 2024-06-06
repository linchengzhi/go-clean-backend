package mysql

import (
	"context"
	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"gorm.io/gorm"
)

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return ArticleRepo{db: db}
}

func (rp *ArticleRepo) WithTx(tx *gorm.DB) *ArticleRepo {
	return &ArticleRepo{db: tx}
}

func (rp *ArticleRepo) GetById(ctx context.Context, id int) (*entity.Article, error) {
	var article *entity.Article
	err := rp.db.First(&article, id).Error
	return article, err
}

func (rp *ArticleRepo) ListByTitle(ctx context.Context, title string, page, pageSize int) ([]*entity.Article, error) {
	var articles = make([]*entity.Article, 0)
	db := rp.db.Model(&entity.Article{})
	if title != "" {
		db = db.Where("title like ?", "%"+title+"%")
	}
	db.Order("id desc")
	if page > 0 && pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Scan(&articles).Error
	return articles, err
}

func (rp *ArticleRepo) Create(ctx context.Context, article *entity.Article) error {
	err := rp.db.Create(article).Error
	return err
}

func (rp *ArticleRepo) Update(ctx context.Context, article *entity.Article) error {
	err := rp.db.Save(article).Error
	return err
}

func (rp *ArticleRepo) Delete(ctx context.Context, id int) error {
	err := rp.db.Delete(&entity.Article{}, id).Error
	return err
}

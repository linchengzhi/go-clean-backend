package mysql

import (
	"gorm.io/gorm"
)

type RepoMysql struct {
	AccountRepo
	UserRepo
	ArticleRepo
}

func NewRepoMysql(db *gorm.DB) RepoMysql {
	repo := new(RepoMysql)
	repo.AccountRepo = NewAccountRepo(db)
	repo.UserRepo = NewUserRepo(db)
	repo.ArticleRepo = NewArticleRepo(db)
	return *repo
}

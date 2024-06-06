package usecase

import (
	"github.com/linchengzhi/go-clean-backend/Infra/database"
	"github.com/linchengzhi/go-clean-backend/repository/mysql"
)

type UcAll struct {
	IAccountUc
	IArticleUc
	IUserUc
}

func NewUcAll(db *database.Gorm, repoMysql *mysql.RepoMysql) UcAll {
	uc := new(UcAll)
	uc.IAccountUc = NewAccountUsecase(db, repoMysql.AccountRepo, repoMysql.UserRepo)
	uc.IArticleUc = NewArticleUc(repoMysql.ArticleRepo)
	uc.IUserUc = NewUserUsecase(repoMysql.UserRepo)
	return *uc
}

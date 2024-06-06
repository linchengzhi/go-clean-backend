package usecase

import (
	"context"

	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"github.com/linchengzhi/go-clean-backend/repository/mysql"
)

type IUserUc interface {
	// 获取用户信息
	GetUserInfo(ctx context.Context, id int) (*entity.User, error)
	// 更新用户信息
	UpdateUsername(ctx context.Context, id int, username string) (*entity.User, error)
}

type UserUc struct {
	userRepo mysql.UserRepo
}

func NewUserUsecase(userRepo mysql.UserRepo) IUserUc {
	return &UserUc{
		userRepo: userRepo,
	}
}

func (uc *UserUc) GetUserInfo(ctx context.Context, id int) (*entity.User, error) {
	user, err := uc.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUc) UpdateUsername(ctx context.Context, id int, username string) (*entity.User, error) {
	user, err := uc.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	user.Name = username
	err = uc.userRepo.Save(ctx, user)
	return user, nil
}

package usecase

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/linchengzhi/go-clean-backend/domain/cerror"
	"io"
	rand2 "math/rand"
	"strconv"
	"time"

	"github.com/linchengzhi/go-clean-backend/Infra/database"
	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"github.com/linchengzhi/go-clean-backend/repository/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IAccountUc interface {
	// Register 注册
	Register(ctx context.Context, email, password string) error
	// Login 登录
	Login(ctx context.Context, email, password string) (*entity.User, error)
	// Logout 登出
	Logout(ctx context.Context, accountId int) error
}

type AccountUc struct {
	accountRepo mysql.AccountRepo
	userRepo    mysql.UserRepo
	db          *database.Gorm
}

func NewAccountUsecase(db *database.Gorm, accountRepo mysql.AccountRepo, userRepo mysql.UserRepo) IAccountUc {
	return &AccountUc{
		accountRepo: accountRepo,
		userRepo:    userRepo,
		db:          db,
	}
}

func (uc *AccountUc) Register(ctx context.Context, email, password string) error {
	// Check if email exists
	account, err := uc.accountRepo.GetByEmail(ctx, email)
	if err == nil && account != nil {
		return cerror.ErrEmailExist.WithErr(err)
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return cerror.ErrBusy.WithErr(err)
	}

	account = uc.prepareAccount(email, password)
	user := uc.prepareUser()

	err = uc.db.Transaction(func(tx *gorm.DB) error {
		if err = uc.accountRepo.WithTx(tx).Create(ctx, account); err != nil {
			return cerror.ErrBusy.WithErr(err)
		}

		user.AccountId = account.Id
		if err = uc.userRepo.WithTx(tx).Create(ctx, user); err != nil {
			return cerror.ErrBusy.WithErr(err)
		}
		return nil
	})

	return err
}

func (uc *AccountUc) Login(ctx context.Context, email, password string) (*entity.User, error) {
	account, err := uc.accountRepo.GetByEmail(ctx, email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cerror.ErrNotAccount.WithErr(err)
	}
	if err != nil {
		return nil, cerror.ErrBusy.WithErr(err)
	}
	if account.Password != getPassword(password, account.Salt) {
		return nil, cerror.ErrPassword
	}
	user, err := uc.userRepo.GetByAccountId(ctx, account.Id)
	if err != nil {
		return nil, cerror.ErrBusy.WithErr(err)
	}
	return user, nil
}

func (uc *AccountUc) Logout(ctx context.Context, accountId int) error {
	return nil
}

func (uc *AccountUc) prepareAccount(email, password string) *entity.Account {
	salt := generateSalt()
	hashedPassword := getPassword(password, salt)
	now := time.Now().UTC()

	return &entity.Account{
		Email:      email,
		Salt:       salt,
		Password:   hashedPassword,
		CreateTime: now,
		UpdateTime: now,
	}
}

func (uc *AccountUc) prepareUser() *entity.User {
	randomSuffix := strconv.FormatInt(rand2.Int63n(1000), 10)
	now := time.Now().UTC()

	return &entity.User{
		Name:       "test" + randomSuffix,
		Birthday:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Profile:    "so lazy",
		CreateTime: now,
		UpdateTime: now,
	}
}

func generateSalt() string {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

func getPassword(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	var res []byte
	res = h.Sum(nil)
	for i := 0; i < 3; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}

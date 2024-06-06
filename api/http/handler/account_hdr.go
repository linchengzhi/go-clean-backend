package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/api/http/middleware"
	"github.com/linchengzhi/go-clean-backend/usecase"
	"github.com/linchengzhi/go-clean-backend/util"
	"go.uber.org/zap"
)

type AccountHdr struct {
	accountUc usecase.IAccountUc
	log       *zap.Logger
}

func NewAccountHandler(uc usecase.IAccountUc, log *zap.Logger) *AccountHdr {
	return &AccountHdr{
		uc,
		log,
	}
}

func (hdr *AccountHdr) Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	err := hdr.accountUc.Register(c, email, password)
	if err != nil {
		hdr.log.Error("Register failed", zap.String("email", email), zap.Error(err))
		util.RespondErr(c, err)
		return
	}
	hdr.log.Sugar().Debugf("Register success, email: %s", email)
	util.Respond(c, nil)
}

// Login 登录
func (hdr *AccountHdr) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user, err := hdr.accountUc.Login(c, email, password)
	if err != nil {
		util.RespondErr(c, err)
		return
	}
	middleware.SetSession(c, user)
	hdr.log.Debug("Login success", zap.Any("user", user))
	util.Respond(c, user)
}

// Logout 登出
func (hdr *AccountHdr) Logout(c *gin.Context) {
	userId := c.GetInt("user_id")
	err := hdr.accountUc.Logout(c, userId)
	if err != nil {
		util.RespondErr(c, err)
		return
	}
	middleware.DelSession(c)
	hdr.log.Debug("Logout success", zap.Int("user_id", userId))
	util.Respond(c, nil)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/api/http/handler"
	"github.com/linchengzhi/go-clean-backend/usecase"
	"go.uber.org/zap"
)

func NewAccountRouter(uc usecase.UcAll, log *zap.Logger, public *gin.RouterGroup, protected *gin.RouterGroup) {
	ad := handler.NewAccountHandler(uc.IAccountUc, log)

	public.POST("register", ad.Register)
	public.POST("login", ad.Login)

	protected.POST("logout", ad.Logout)
}

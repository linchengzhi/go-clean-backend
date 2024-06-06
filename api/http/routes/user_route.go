package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/api/http/handler"
	"github.com/linchengzhi/go-clean-backend/usecase"
	"go.uber.org/zap"
)

func NewUserRouter(uc usecase.UcAll, log *zap.Logger, protected *gin.RouterGroup) {
	ud := handler.NewUserHandler(uc.IUserUc)

	pt := protected.Group("user")

	pt.GET("get", ud.Get)
	pt.POST("update/name", ud.UpdateName)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/api/http/handler"
	"github.com/linchengzhi/go-clean-backend/usecase"
	"go.uber.org/zap"
)

func NewArticleRouter(uc usecase.UcAll, log *zap.Logger, public *gin.RouterGroup, protected *gin.RouterGroup) {
	ud := handler.NewArticleHandler(uc.IArticleUc, log)

	pu := public.Group("article")
	pu.GET("get", ud.Get)
	pu.GET("list", ud.List)

	pt := protected.Group("article")
	pt.POST("add", ud.Add)
}

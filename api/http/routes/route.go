package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/api/http/middleware"
	"github.com/linchengzhi/go-clean-backend/usecase"
	"go.uber.org/zap"
)

func SetRoutes(uc usecase.UcAll, log *zap.Logger, gin *gin.Engine) {
	// All Public APIs
	publicRouter := gin.Group("")

	// All Private APIs
	protectedRouter := gin.Group("")

	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.CheckLogin())

	NewAccountRouter(uc, log, publicRouter, protectedRouter)
	NewUserRouter(uc, log, protectedRouter)
	NewArticleRouter(uc, log, publicRouter, protectedRouter)
}

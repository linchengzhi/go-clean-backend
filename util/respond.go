package util

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/domain/cerror"
	"net/http"
)

func Respond(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func RespondErr(c *gin.Context, err error) {
	customErr, ok := err.(*cerror.CustomError)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": customErr.GetCode(),
			"msg":  customErr.GetMsg(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": cerror.ErrSystem.GetCode(),
			"msg":  cerror.ErrSystem.GetMsg(),
		})
	}
}

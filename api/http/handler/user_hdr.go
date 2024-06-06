package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/usecase"
)

type UserHdr struct {
	userUc usecase.IUserUc
}

func NewUserHandler(uc usecase.IUserUc) *UserHdr {
	return &UserHdr{
		uc,
	}
}

func (hdr *UserHdr) Get(c *gin.Context) {
	userId := c.GetInt("user_id")
	user, err := hdr.userUc.GetUserInfo(c, userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    user,
	})
}

func (hdr *UserHdr) UpdateName(c *gin.Context) {
	userId := c.GetInt("user_id")
	name := c.PostForm("name")
	user, err := hdr.userUc.UpdateUsername(c, userId, name)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    user,
	})
}

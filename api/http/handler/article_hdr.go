package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/usecase"
	"github.com/linchengzhi/goany"
	"go.uber.org/zap"
)

type ArticleHdr struct {
	articleUc usecase.IArticleUc
	log       *zap.Logger
}

func NewArticleHandler(uc usecase.IArticleUc, log *zap.Logger) *ArticleHdr {
	return &ArticleHdr{
		uc,
		log,
	}
}

// Add Article
func (hdr *ArticleHdr) Add(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	err := hdr.articleUc.Add(c, title, content)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// Get Article
func (hdr *ArticleHdr) Get(c *gin.Context) {
	id := c.Query("id")
	article, err := hdr.articleUc.GetByID(c, goany.ToInt(id))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    article,
	})
}

// List Article
func (hdr *ArticleHdr) List(c *gin.Context) {
	title := c.Query("title")
	page := c.Query("page")
	pageSize := c.Query("page_size")
	articles, err := hdr.articleUc.List(c, title, goany.ToInt(page), goany.ToInt(pageSize))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    articles,
	})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/services"
	"net/http"
)

type ArticleCtlParam struct {
	BaseCtlParams
	TagService *services.TagService
}

type ArticleController struct {
	BaseController
	tagService *services.TagService
}

func NewArticleController(t ArticleCtlParam) *ArticleController {
	return &ArticleController {
		tagService: t.TagService,
		BaseController: BaseController {
			logs:t.Logs,
		},
	}
}

func (t *ArticleController) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "article"})
}

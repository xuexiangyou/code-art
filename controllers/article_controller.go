package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/interfaces"
	"gorm.io/gorm"
	"net/http"
)

type ArticleCtlParam struct {
	BaseCtlParams
}

type ArticleController struct {
	BaseController
}

var _ interfaces.ArticleController = ArticleController{}

func NewArticleController(t ArticleCtlParam) ArticleController {
	return ArticleController{
		BaseController: NewBaseController(t.Db, t.Redis),
	}
}

func (t ArticleController) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "article"})
}

func (t ArticleController) CreateArticle(c *gin.Context) {
	var createArticleParam forms.CreateArticle
	err := c.ShouldBind(&createArticleParam)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS, forms.Translate(err))
		return
	}

	//获取开启的事物
	txHandle := c.MustGet("db_trx").(*gorm.DB)
	articleService := t.WithTrxDb(txHandle).newArticleService()
	articleRet, err := articleService.CreateArticle(createArticleParam)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS)
		return
	}

	common.WrapContext(c).Success(articleRet)
}

func (t ArticleController) UpdateArticle(c *gin.Context) {
	var updateArticleParam forms.UpdateArticle
	err := c.ShouldBind(&updateArticleParam)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS, forms.Translate(err))
		return
	}

	articleService := t.newArticleService()
	err = articleService.UpdateArticle(updateArticleParam)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS, err.Error())
		return
	}

	common.WrapContext(c).Success("")
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/controllers"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/interfaces"
	"github.com/xuexiangyou/code-art/middleware"
	"gorm.io/gorm"
	"net/http"
)

type ArticleController struct {
	controllers.BaseController
}

var _ interfaces.ArticleController = ArticleController{}

func newArticleRoutes(handler *gin.RouterGroup, f common.FxCommonParams) {
	r := newArticleController(f)
	h := handler.Group("/article")
	{
		h.GET("/get", r.GetArticle)
		h.POST("/create", middleware.DBTransactionMiddleware(f.Db), r.CreateArticle)
		h.POST("/update", r.UpdateArticle)
	}
}

func newArticleController(f common.FxCommonParams) ArticleController {
	return ArticleController{
		BaseController: controllers.NewBaseController(f.Db, f.Redis),
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
	articleService := t.WithTrxDb(txHandle).NewArticleService()
	articleRet, err := articleService.CreateArticle(createArticleParam)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS, err.Error())
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

	articleService := t.NewArticleService()
	err = articleService.UpdateArticle(updateArticleParam)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS, err.Error())
		return
	}

	common.WrapContext(c).Success("")
}

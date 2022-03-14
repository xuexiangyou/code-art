package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/services"
	"gorm.io/gorm"
	"net/http"
)

type ArticleCtlParam struct {
	BaseCtlParams
	TagService     *services.TagService
	ArticleService *services.ArticleService
}

type ArticleController struct {
	BaseController
	tagService     *services.TagService
	articleService *services.ArticleService
}

func NewArticleController(t ArticleCtlParam) *ArticleController {
	return &ArticleController{
		tagService:     t.TagService,
		articleService: t.ArticleService,
		BaseController: BaseController{
			logs: t.Logs,
		},
	}
}

func (t *ArticleController) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "article"})
}

func (t *ArticleController) CreateArticle(c *gin.Context) {
	appG := common.Gin{C: c}

	var createArticleParam forms.CreateArticle
	err := c.ShouldBind(&createArticleParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": forms.Translate(err)})
		return
	}

	//获取开启的事物
	txHandle := c.MustGet("db_trx").(*gorm.DB)

	//插入tag表数据
	tagData := &entity.Tag{
		Name: createArticleParam.Name,
	}
	tagRet, err := t.tagService.WithTrx(txHandle).CreateTag(tagData)
	if err != nil {
		t.logs.AppLog.Error("插入tag数据失败", err)
		appG.Response(http.StatusInternalServerError, common.INVALID_PARAMS, nil)
		return
	}

	fmt.Println("tag_ret---", tagRet)

	//插入文章数据
	articleData := &entity.Article{
		TagId: tagRet.Id,
		Title: createArticleParam.Title,
	}
	articleRet, err := t.articleService.WithThr(txHandle).CreateArticle(articleData)
	if err != nil {
		t.logs.AppLog.Error("插入tag数据失败", err)
		appG.Response(http.StatusInternalServerError, common.INVALID_PARAMS, nil)
	}
	fmt.Println("article_ret---", articleRet)

	appG.Response(http.StatusOK, common.SUCCESS, "ok")
}

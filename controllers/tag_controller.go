package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/services"
	"net/http"
	"time"
)

type TagCtlParam struct {
	BaseCtlParams
	TagService *services.TagService
}

type TagController struct {
	BaseController
	tagService *services.TagService
}

func NewTagController(t TagCtlParam) *TagController {
	return &TagController {
		tagService: t.TagService,
		BaseController: BaseController {
			logs:t.Logs,
		},
	}
}

type TestData struct {
	Name string `json:"name"`
}

func (t *TagController) TestTag(c *gin.Context) {
	common.WrapContext(c).Success("222")
}

func (t *TagController) GetTag(c *gin.Context) {

	time.Sleep(5 * time.Second)

	//打印info日志
	t.logs.AppLog.Info("hahahh")
	t.logs.AppLog.Error("哈哈哈哈哈")

	var getTagParam forms.GetTag
	err := c.ShouldBindQuery(&getTagParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": forms.Translate(err)})
		return
	}

	ret, err := t.tagService.GetById(getTagParam.Id)
	if err != nil {
		fmt.Println("------", err)
		c.JSON(http.StatusBadRequest, "get tag data invalid")
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (t *TagController) UpdateTag(c *gin.Context) {
	var updateTagParam forms.UpdateTag
	//json 参数绑定
	err := c.ShouldBindJSON(&updateTagParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": forms.Translate(err)})
		return
	}

	c.JSON(http.StatusOK, updateTagParam)
}

func (t *TagController) CreateTag(c *gin.Context) {
	var createTagParam forms.CreateTag
	//json 参数绑定
	err := c.ShouldBindJSON(&createTagParam)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS)
		return
	}

	tagData := &entity.Tag{
		Name: createTagParam.Name,
	}
	tagRet, err := t.tagService.CreateTag(tagData)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS)
		return
	}
	common.WrapContext(c).Success(tagRet)
}

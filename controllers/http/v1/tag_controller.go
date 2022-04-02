package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/interfaces"
	"net/http"
)

type TagCtlParam struct {
	BaseCtlParams
}

type TagController struct {
	BaseController
}

var _ interfaces.TagController = TagController{}

func newTagRoutes(handler *gin.RouterGroup, f common.FxCommonParams) {
	r := newTagController(f)
	h := handler.Group("/tag")
	{
		h.GET("/get", r.GetTag)
		h.GET("/test", r.TestTag)
		h.POST("/create", r.CreateTag)
		h.POST("/update", r.UpdateTag)
	}
}

func newTagController(f common.FxCommonParams) TagController {
	return TagController{
		BaseController: NewBaseController(f.Db, f.Redis),
	}
}

func (t TagController) TestTag(c *gin.Context) {
	log := c.MustGet("logger").(*logrus.Entry)
	log.Info("HAHAHAHHAHAHAHAH")

	common.WrapContext(c).Success("222")
}

func (t TagController) GetTag(c *gin.Context) {
	log := c.MustGet("logger").(*logrus.Entry)
	log.Info("eeeeeeee")

	var getTagParam forms.GetTag
	err := c.ShouldBindQuery(&getTagParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": forms.Translate(err)})
		return
	}

	ret, err := t.newTagService().GetById(getTagParam.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "get tag data invalid")
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (t TagController) UpdateTag(c *gin.Context) {
	var updateTagParam forms.UpdateTag
	//json 参数绑定
	err := c.ShouldBindJSON(&updateTagParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": forms.Translate(err)})
		return
	}

	c.JSON(http.StatusOK, updateTagParam)
}

func (t TagController) CreateTag(c *gin.Context) {
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
	tagRet, err := t.newTagService().CreateTag(tagData)
	if err != nil {
		common.WrapContext(c).Error(http.StatusInternalServerError, common.INVALID_PARAMS)
		return
	}
	common.WrapContext(c).Success(tagRet)
}

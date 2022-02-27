package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/services"
	"net/http"
)

type TagController struct {
	tagService *services.TagService
}

func NewTagController(tag *services.TagService) *TagController {
	return &TagController{
		tagService: tag,
	}
}

func (t *TagController) GetTag(c *gin.Context) {
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
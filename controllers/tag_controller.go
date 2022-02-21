package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/services"
	"net/http"
	"strconv"
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
	tagId, err := strconv.ParseInt(c.Query("tag_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}
	ret, err := t.tagService.GetById(tagId)
	if err != nil {
		fmt.Println("------", err)
		c.JSON(http.StatusBadRequest, "get tag data invalid")
		return
	}

	c.JSON(http.StatusOK, ret)
}
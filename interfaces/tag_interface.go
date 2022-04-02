package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/domain/entity"
)

type TagController interface {
	TestTag(c *gin.Context)
	GetTag(c *gin.Context)
	UpdateTag(c *gin.Context)
	CreateTag(c *gin.Context)
}

type TagStrategy interface {
	GetTag(int64) (*entity.Tag, error)
	ListByIds([]int64) ([]*entity.Tag, error)
	CreateTag(*entity.Tag) (*entity.Tag, error)
}

type TagRepo interface {
	GetTag(id int64) (*entity.Tag, error)
	ListById([]int64) ([]*entity.Tag, error)
	CreateTag(tag *entity.Tag) (*entity.Tag, error)
}

type TagCache interface {
	GetById(id int64) (string, error)
}

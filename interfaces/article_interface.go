package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/domain/entity"
)

type ArticleController interface {
	GetArticle(c *gin.Context)
	CreateArticle(c *gin.Context)
}

type ArticleStrategy interface {
	CreateArticle(*entity.Article) (*entity.Article, error)
	UpdateTitleById(id int64, title string) error
}

type ArticleRepo interface {
	CreateArticle(*entity.Article) (*entity.Article, error)
	UpdateTitleById(id int64, title string) error
}

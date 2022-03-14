package strategy

import (
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/domain/storage"
	"gorm.io/gorm"
)

type ArticleStrategy struct {
	db    *gorm.DB
	cache *redis.Client
}

type ArticleRepository interface {
	CreateArticle(*entity.Article) (*entity.Article, error)
	WithTrx(db *gorm.DB) *ArticleStrategy
}

var _ ArticleRepository = &ArticleStrategy{}

func NewArticleStrategy(db *gorm.DB, cache *redis.Client) *ArticleStrategy {
	return &ArticleStrategy{db, cache}
}

//WithTrx 设置事物数据库链接
func (a ArticleStrategy) WithTrx(db *gorm.DB) *ArticleStrategy {
	a.db = db
	return &a
}

//getArticleModel 获取tag数据库对象
func (a *ArticleStrategy) getArticleModel() *storage.ArticleModel {
	return storage.NewArticleModel(a.db)
}

//CreateArticle 创建文章
func (a *ArticleStrategy) CreateArticle(article *entity.Article) (*entity.Article, error) {
	return a.getArticleModel().CreateArticle(article)
}


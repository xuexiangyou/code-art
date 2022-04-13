package controllers

import (
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/config"
	"github.com/xuexiangyou/code-art/domain/cache"
	"github.com/xuexiangyou/code-art/domain/storage"
	"github.com/xuexiangyou/code-art/domain/strategy"
	"github.com/xuexiangyou/code-art/services"
	"gorm.io/gorm"
)

type BaseController struct {
	Config *config.Config
	Db     *gorm.DB
	Redis  *redis.Client
}

func NewBaseController(config * config.Config, db *gorm.DB, redis *redis.Client) BaseController {
	return BaseController{
		Config: config,
		Db:     db,
		Redis:  redis,
	}
}

//WithTrxDb 设置事物的db
func (b BaseController) WithTrxDb(db *gorm.DB) BaseController {
	b.Db = db
	return b
}

func (b BaseController) NewTagService() services.TagService {
	tagStrategy := b.NewTagStrategy()
	return services.NewTagService(tagStrategy)
}

func (b BaseController) NewArticleStrategy() strategy.ArticleStrategy {
	articleModel := storage.NewArticleModel(b.Db)
	return strategy.NewArticleStrategy(articleModel)
}

func (b BaseController) NewTagStrategy() strategy.TagStrategy {
	tagModel := storage.NewTagModel(b.Db)
	tagCache := cache.NewTagCache(b.Redis)
	return strategy.NewTagStrategy(tagModel, tagCache)
}

func (b BaseController) NewArticleService() services.ArticleService {
	articleStrategy := b.NewArticleStrategy()
	tagStrategy := b.NewTagStrategy()
	return services.NewArticleService(articleStrategy, tagStrategy)
}

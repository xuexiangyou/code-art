package v1

import (
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/domain/cache"
	"github.com/xuexiangyou/code-art/domain/storage"
	"github.com/xuexiangyou/code-art/domain/strategy"
	"github.com/xuexiangyou/code-art/services"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type BaseCtlParams struct {
	fx.In
	Db    *gorm.DB
	Redis *redis.Client
}

type BaseController struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewBaseController(db *gorm.DB, redis *redis.Client) BaseController {
	return BaseController{
		db:    db,
		redis: redis,
	}
}

//WithTrxDb 设置事物的db
func (b BaseController) WithTrxDb(db *gorm.DB) BaseController {
	b.db = db
	return b
}

func (b BaseController) newTagService() services.TagService {
	tagStrategy := b.newTagStrategy()
	return services.NewTagService(tagStrategy)
}


func (b BaseController) newArticleStrategy() strategy.ArticleStrategy {
	articleModel := storage.NewArticleModel(b.db)
	return strategy.NewArticleStrategy(articleModel)
}

func (b BaseController) newTagStrategy() strategy.TagStrategy {
	tagModel := storage.NewTagModel(b.db)
	tagCache := cache.NewTagCache(b.redis)
	return strategy.NewTagStrategy(tagModel, tagCache)
}

func (b BaseController) newArticleService() services.ArticleService {
	articleStrategy := b.newArticleStrategy()
	tagStrategy := b.newTagStrategy()
	return services.NewArticleService(articleStrategy, tagStrategy)
}


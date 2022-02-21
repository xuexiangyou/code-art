package strategy

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/domain/cache"
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/domain/repository"
	"github.com/xuexiangyou/code-art/domain/storage"
	"gorm.io/gorm"
)

type TagStrategy struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewTagStrategy(db *gorm.DB, cache *redis.Client) *TagStrategy {
	return &TagStrategy{db, cache}
}

var _ repository.TagRepository = &TagStrategy{}

//getTagCache 获取tag缓存对象
func (t *TagStrategy) getTagCache() *cache.TagCache {
	return cache.NewTagCache(t.cache)
}

//getTagModel 获取tag数据库对象
func (t *TagStrategy) getTagModel() *storage.TagModel {
	return storage.NewTagModel(t.db)
}

//GetTag 根据id获取tag
func (t *TagStrategy) GetTag(id int64) (*entity.Tag, error) {
	var err error

	tagCacheData, err := t.getTagCache().GetById(id)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	//如果缓存不存在数据则读取数据库中数据
	if tagCacheData == "" {
		ret, err := t.getTagModel().GetTag(id)
		if err != nil {
			return nil, err
		}
		return ret, nil
	} else {
		var ret entity.Tag
		if tagCacheData != "" {
			err = json.Unmarshal([]byte(tagCacheData), &ret)
		}
		return &ret, err
	}
}

//ListByIds 根据ids数组获取tag列表
func (t *TagStrategy) ListByIds(ids []int64) ([]*entity.Tag, error) {
	return nil, nil
}

package strategy

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/interfaces"
)

type TagStrategy struct {
	tagRepo interfaces.TagRepo
	tagCache interfaces.TagCache
}

func NewTagStrategy(tagRepo interfaces.TagRepo, tagCache interfaces.TagCache) TagStrategy {
	return TagStrategy{
		tagRepo: tagRepo,
		tagCache: tagCache,
	}
}

//GetTag 根据id获取tag
func (t TagStrategy) GetTag(id int64) (*entity.Tag, error) {
	var err error

	tagCacheData, err := t.tagCache.GetById(id)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	//如果缓存不存在数据则读取数据库中数据
	if tagCacheData == "" {
		ret, err := t.tagRepo.GetTag(id)
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

//CreateTag 创建tag表记录
func (t TagStrategy) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	return t.tagRepo.CreateTag(tag)
}

//ListByIds 根据ids数组获取tag列表
func (t TagStrategy) ListByIds(ids []int64) ([]*entity.Tag, error) {
	return nil, nil
}

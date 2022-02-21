package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type TagCache struct {
	cache *redis.Client
}

var (
	TagDetailStr = "tag_detail_str_%d"
)

func NewTagCache(redis *redis.Client) *TagCache {
	return &TagCache{
		cache: redis,
	}
}

func (c *TagCache) getTagDetailStrKey(id int64) string {
	return fmt.Sprintf(TagDetailStr, id)
}

func (c *TagCache) GetById(id int64) (string, error) {
	key := c.getTagDetailStrKey(id)
	return c.cache.Get(context.Background(), key).Result()
}

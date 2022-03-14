package strategy

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BaseStrategy struct {
	db    *gorm.DB
	cache *redis.Client
}

func (b BaseStrategy) WithTrx(db *gorm.DB) *BaseStrategy {
	b.db = db
	return &b
}

package services

import (
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/domain/strategy"
	"gorm.io/gorm"
)

type TagService struct {
	tag *strategy.TagStrategy
}

func NewTagService(tag *strategy.TagStrategy) *TagService {
	return &TagService{tag: tag}
}

func (t TagService) WithTrx(db *gorm.DB) *TagService {
	t.tag = t.tag.WithTrx(db)
	return &t
}

func (t *TagService) GetById(tagId int64) (*entity.Tag, error) {
	ret, err := t.tag.GetTag(tagId)
	return ret, err
}

func (t *TagService) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	ret, err := t.tag.CreateTag(tag)
	return ret, err
}
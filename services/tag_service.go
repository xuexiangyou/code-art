package services

import (
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/domain/strategy"
)

type TagService struct {
	tag *strategy.TagStrategy
}

func NewTagService(tag *strategy.TagStrategy) *TagService {
	return &TagService{tag: tag}
}

func (t *TagService) GetById(tagId int64) (*entity.Tag, error) {
	ret, err := t.tag.GetTag(tagId)
	return ret, err
}
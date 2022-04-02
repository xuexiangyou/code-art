package services

import (
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/interfaces"
)

type TagService struct {
	tag interfaces.TagStrategy
}

func NewTagService(tag interfaces.TagStrategy) TagService {
	return TagService{tag: tag}
}

func (t TagService) GetById(tagId int64) (*entity.Tag, error) {
	ret, err := t.tag.GetTag(tagId)
	return ret, err
}

func (t TagService) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	ret, err := t.tag.CreateTag(tag)
	return ret, err
}
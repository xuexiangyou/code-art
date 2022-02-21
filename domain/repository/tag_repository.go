package repository

import "github.com/xuexiangyou/code-art/domain/entity"

type TagRepository interface {
	GetTag(int64) (*entity.Tag, error)
	ListByIds([]int64) ([]*entity.Tag, error)
}

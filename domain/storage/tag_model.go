package storage

import (
	"github.com/xuexiangyou/code-art/domain/entity"
	"gorm.io/gorm"
)

type TagModel struct {
	db *gorm.DB
}

func NewTagModel(db *gorm.DB) *TagModel {
	return &TagModel{db}
}

func (t *TagModel) TableName() string {
	return "test_tag"
}

func (t *TagModel) GetTag(id int64) (*entity.Tag, error) {
	var tag entity.Tag
	result := t.db.Table(t.TableName()).First(&tag, id)
	return &tag, result.Error
}

func (t *TagModel) ListById([]int64) ([]*entity.Tag, error) {
	return nil, nil
}


package storage

import (
	"fmt"
	"github.com/xuexiangyou/code-art/domain/entity"
	"gorm.io/gorm"
)

type ArticleModel struct {
	db *gorm.DB
}

func NewArticleModel(db *gorm.DB) ArticleModel {
	return ArticleModel{db}
}

func (a ArticleModel) TableName() string {
	return "test_article"
}

func (a ArticleModel) CreateArticle(article *entity.Article) (*entity.Article, error) {
	err := a.db.Table(a.TableName()).Create(article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a ArticleModel) UpdateTitleById(id int64, title string) error {
	ret := a.db.Table(a.TableName()).Where("id = ?",  id).Update("title", title)
	fmt.Println("------", ret.RowsAffected)

	return ret.Error
}

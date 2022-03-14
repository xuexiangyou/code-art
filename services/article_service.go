package services

import (
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/domain/strategy"
	"gorm.io/gorm"
)

type ArticleService struct {
	article *strategy.ArticleStrategy
}

func NewArticleService(article *strategy.ArticleStrategy) *ArticleService {
	return &ArticleService{article}
}

func (a ArticleService) WithThr(db *gorm.DB) *ArticleService {
	a.article = a.article.WithTrx(db)
	return &a
}

func (a *ArticleService) CreateArticle(article *entity.Article) (*entity.Article, error) {
	ret, err := a.article.CreateArticle(article)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

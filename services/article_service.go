package services

import (
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/interfaces"
)

type ArticleService struct {
	article interfaces.ArticleStrategy
	tag     interfaces.TagStrategy
}

func NewArticleService(article interfaces.ArticleStrategy, tag interfaces.TagStrategy) ArticleService {
	return ArticleService{
		article: article,
		tag: tag,
	}
}

func (a *ArticleService) CreateArticle(createArticleParam forms.CreateArticle) (*entity.Article, error) {
	tagData := &entity.Tag{
		Name: createArticleParam.Name,
	}
	tagRet, err := a.tag.CreateTag(tagData)
	if err != nil {
		return nil, err
	}

	articleData := &entity.Article{
		TagId: tagRet.Id,
		Title: createArticleParam.Title,
	}
	articleRet, err := a.article.CreateArticle(articleData)
	if err != nil {
		return nil, err
	}
	return articleRet, nil
}

func (a *ArticleService) UpdateArticle(updateArticleParam forms.UpdateArticle) error {
	err := a.article.UpdateTitleById(updateArticleParam.Id, updateArticleParam.Title)
	return err
}

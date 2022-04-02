package strategy

import (
	"github.com/xuexiangyou/code-art/domain/entity"
	"github.com/xuexiangyou/code-art/interfaces"
)

type ArticleStrategy struct {
	articleRepo interfaces.ArticleRepo
}

func NewArticleStrategy(articleRepo interfaces.ArticleRepo) ArticleStrategy {
	return ArticleStrategy{
		articleRepo: articleRepo,
	}
}

//CreateArticle 创建文章
func (a ArticleStrategy) CreateArticle(article *entity.Article) (*entity.Article, error) {
	return a.articleRepo.CreateArticle(article)
}

func (a ArticleStrategy) UpdateTitleById(id int64, title string) error {
	return a.articleRepo.UpdateTitleById(id, title)
}


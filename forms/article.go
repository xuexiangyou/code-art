package forms

type CreateArticle struct {
	Name 		string 	`json:"name" binding:"required"`
	Title		string	`json:"title" binding:"required"`
}

type UpdateArticle struct {
	Id 			int64	`json:"id" binding:"required"`
	Title		string	`json:"title" binding:"required"`
}

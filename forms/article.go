package forms

type CreateArticle struct {
	Name 		string 	`json:"name" binding:"required"`
	Title		string	`json:"title" binding:"required"`
}

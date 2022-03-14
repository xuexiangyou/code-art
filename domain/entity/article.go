package entity

type Article struct {
	Id  	int64	`json:"id"`
	TagId   int64	`json:"tag_id"`
	Title	string	`json:"title"`
}

func (a *Article) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	if a.Id == 0 {
		errorMessages["id_required"] = "id is required"
	}
	if a.TagId == 0 {
		errorMessages["tagId_required"]= "tagId is required"
	}
	if a.Title == "" {
		errorMessages["title_required"] = "title is required"
	}
	return errorMessages
}

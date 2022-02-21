package entity

type Tag struct {
	Id 		int64 		`json:"id"`
	Name	string		`json:"name"`
}

//Validate 添加校验逻辑
func (t *Tag) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	if t.Id == 0 {
		errorMessages["id_required"] = "id is required"
	}
	if t.Name == "" {
		errorMessages["name_required"] = "name is required"
	}
	return errorMessages
}
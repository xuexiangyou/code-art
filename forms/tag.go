package forms

// GetTag 获取tag参数
type GetTag struct {
	Id int64 `form:"id" binding:"required"`
}

//UpdateTag 更新tag参数
type UpdateTag struct {
	Id   		int64  `form:"id" binding:"required"`
	Name 		string `json:"name" binding:"required,checkName"`
	CreateTime 	string	`json:"create_time" binding:"required,timing"`
}

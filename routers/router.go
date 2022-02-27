package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xuexiangyou/code-art/controllers"
	"github.com/xuexiangyou/code-art/forms"
	"go.uber.org/dig"
)

type RouterParams struct {
	dig.In
	TagController *controllers.TagController
}


func InitRouter(p RouterParams) *gin.Engine {
	//表单校验
	forms.Init()

	r := gin.Default()

	//Registration verification
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timing", forms.Timing)
		if err != nil {
			fmt.Println("fail")
		}
		err = v.RegisterValidation("checkName", forms.CheckName)
		if err != nil {
			fmt.Println("fail")
		}
	}

	r.GET("/get-tag", p.TagController.GetTag)
	r.POST("/update-tag", p.TagController.UpdateTag)

	return r
}


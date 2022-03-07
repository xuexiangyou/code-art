package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xuexiangyou/code-art/controllers"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/middleware/log"
	"go.uber.org/fx"
)

type RouterParams struct {
	fx.In
	Logs              *log.Logs
	TagController     *controllers.TagController
	ArticleController *controllers.ArticleController
}

func InitRouter(p RouterParams) *gin.Engine {
	//表单校验
	forms.Init()

	r := gin.New()
	//Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	//自定义日志格式
	// Logging to a file. todo 这个是把日志从终端切换到日志文件中
	/*f, _ := os.Create("gin.logs")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))*/

	//自定义日志中间件
	r.Use(log.LoggerToFile(p.Logs))

	//Registration verification
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timing", forms.Timing)
		if err != nil {
			fmt.Println("fail") //todo
		}
		err = v.RegisterValidation("checkName", forms.CheckName)
		if err != nil {
			fmt.Println("fail") //todo
		}
	}

	//tag 相关接口定义
	r.GET("/get-tag", p.TagController.GetTag)
	r.POST("/update-tag", p.TagController.UpdateTag)
	r.GET("/test-tag", p.TagController.TestTag)

	//article 相关接口定义
	r.GET("/get-article", p.ArticleController.GetArticle)

	return r
}

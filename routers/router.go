package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/config"
	v1 "github.com/xuexiangyou/code-art/controllers/http/v1"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/middleware"
	"log"
	"time"
)

//设置gin框架的模式
func setGinMode(c *config.Config) {
	switch c.Env {
	case common.AppEnvDev:
		gin.SetMode(gin.DebugMode)
	case common.AppEnvTest:
		gin.SetMode(gin.TestMode)
	case common.AppEnvPrd:
		gin.SetMode(gin.ReleaseMode)
	}
}

func InitRouter(p common.FxCommonParams) *gin.Engine {
	//表单校验
	forms.Init()

	//设置模式
	setGinMode(p.Config)

	//new engine
	r := gin.New()

	//Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	//设置接口请求日志
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
	}))

	//定义日志transaction中间件
	r.Use(middleware.JsonLogMiddleware())

	//cors
	r.Use(middleware.CORS(middleware.CORSOptions{}))

	//Registration verification
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timing", forms.Timing)
		if err != nil {
			log.Fatal("register timing verify fail", err)
		}
		err = v.RegisterValidation("checkName", forms.CheckName)
		if err != nil {
			log.Fatal("register checkName verify fail", err)
		}
	}

	//设置v1版本的接口路由
	v1.NewRouter(r, p)

	return r
}

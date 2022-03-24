package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/config"
	"github.com/xuexiangyou/code-art/controllers"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/middleware/cors"
	"github.com/xuexiangyou/code-art/middleware/log"
	"github.com/xuexiangyou/code-art/middleware/transaction"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type RouterParams struct {
	fx.In
	Logs              *log.Logs
	TagController     *controllers.TagController
	ArticleController *controllers.ArticleController
	Db                *gorm.DB
	Config 			  *config.Config
}

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

func InitRouter(p RouterParams) *gin.Engine {
	//表单校验
	forms.Init()

	//设置模式
	setGinMode(p.Config)

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
	//r.Use(log.LoggerToFile(p.Logs))

	//定义日志transaction中间件
	r.Use(log.JsonLogMiddleware())

	//cors
	r.Use(cors.CORS(cors.CORSOptions{}))

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
	r.POST("/create-tag", p.TagController.CreateTag)

	//article 相关接口定义
	r.GET("/get-article", p.ArticleController.GetArticle)
	r.POST("/create-article", transaction.DBTransactionMiddleware(p.Db), p.ArticleController.CreateArticle)

	return r
}

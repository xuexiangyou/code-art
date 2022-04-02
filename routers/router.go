package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/config"
	"github.com/xuexiangyou/code-art/controllers"
	"github.com/xuexiangyou/code-art/forms"
	"github.com/xuexiangyou/code-art/middleware"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
)

type RouterParams struct {
	fx.In
	Db                *gorm.DB
	Config            *config.Config
	TagController     controllers.TagController
	ArticleController controllers.ArticleController
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

	//new engine
	r := gin.New()

	//Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

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

	//tag 相关接口定义
	r.GET("/get-tag", p.TagController.GetTag)
	r.POST("/update-tag", p.TagController.UpdateTag)
	r.GET("/test-tag", p.TagController.TestTag)
	r.POST("/create-tag", p.TagController.CreateTag)

	//article 相关接口定义
	r.GET("/get-article", p.ArticleController.GetArticle)
	r.POST("/create-article", middleware.DBTransactionMiddleware(p.Db), p.ArticleController.CreateArticle)
	r.POST("/update-article", p.ArticleController.UpdateArticle)

	return r
}

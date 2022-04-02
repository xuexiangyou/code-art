package forms

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"time"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate = binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

	// 所以这一步注册要放到trans初始化的后面
	if err := validate.RegisterTranslation(
		"timing",
		trans,
		registerTranslator("timing", "{0}必须要晚于当前日期"),
		translate,
	); err != nil {
		return
	}

	// 所以这一步注册要放到trans初始化的后面
	if err := validate.RegisterTranslation(
		"checkName",
		trans,
		registerTranslator("checkName", "{0}必须为llg"),
		translate,
	); err != nil {
		return
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

//Translate 翻译错误信息
func Translate(err error) string {
	var msg string
	errors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, err := range errors {
			msg += err.Translate(trans) + ","
		}
	} else {
		msg = err.Error()
	}

	//for _, err := range errors{
	//	result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	//}
	return msg
}

//Timing 自定义时间检验函数
func Timing(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(string); ok {
		timeT, _ := time.Parse("2006-01-02", date)
		today := time.Now()
		if today.After(timeT) {
			return false
		}
	}
	return true
}

func CheckName(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val != "llg" {
		return false
	}

	return true
}

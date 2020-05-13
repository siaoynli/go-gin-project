//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:47

package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/ini.v1"
	"log"
	"reflect"
	"strings"
)

type Validators struct {
	context *gin.Context
	validate  *validator.Validate
	trans     ut.Translator
	viewModel interface{}
	Errors    map[string]string //所有错误信息
	Error     string  //显示一个校验错误
}

func New() *Validators {
	chinese := zh.New()
	uni := ut.New(chinese)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &Validators{
		validate:  validate,
		trans:     trans,
		Errors:    make(map[string]string),
	}
}

func (v *Validators)Bind(context *gin.Context,viewModel interface{}) (*Validators,error) {

	if err:=context.ShouldBind(&viewModel);err != nil {
		return &Validators{},err
	}
	v.context= context
	v.viewModel=viewModel
	fmt.Println(v.viewModel)
	return v,nil
}


func (v *Validators) Validated() *Validators {
	err := v.validate.Struct(v.viewModel)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		validType := reflect.TypeOf(v.viewModel)
		name := validType.Name()
		for _, e := range errs {
			field, _ := validType.FieldByName(e.Field())
			//反射获取到表单字段
			formField := field.Tag.Get("form")
			//读取语言配置文件
			langFile := "./resources/langs/zh_CN.ini"
			lang, err := ini.Load(langFile)
			if err != nil {
				log.Fatal(err.Error())
			}
			//获取分区配置文件
			langText := lang.Section(name).Key(formField).String()
			tranText := e.Translate(v.trans)
			//如果翻译有内容，则替换
			if len(langText) == 0 {
				//搜索全局配置文件
				langText = lang.Section("").Key(formField).String()
			}
			if len(langText) > 0 {
				tranText = strings.Replace(tranText, e.Field(), langText, -1)
			}
			v.Errors[formField] = tranText

			if len(v.Error) ==0 {
				v.Error = tranText
			}
		}
	}
	return v
}

//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 17:22

package Request

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/ini.v1"
	"log"
	"reflect"
	"strings"
)

type Validator struct {
	validate       *validator.Validate
	trans          ut.Translator
	validatorModel interface{}       //validator struct
	Errors         map[string]string //所有错误信息
	Error          string            //显示一个校验错误
	cfg            *ini.File
}

func NewValidator() *Validator {
	chinese := zh.New()
	uni := ut.New(chinese)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatalln(err.Error())
	}
	langFile := "./Resources/langs/zh_CN.ini"
	cfg, err := ini.Load(langFile)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Validator{
		validate: validate,
		trans:    trans,
		cfg:      cfg,
		Errors:   make(map[string]string),
	}
}

func (v *Validator) Validated(validatorModel interface{}) *Validator {
	v.validatorModel = validatorModel
	err := v.validate.Struct(v.validatorModel)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		validType := reflect.TypeOf(v.validatorModel)
		name := validType.Name()
		for _, e := range errs {
			field, _ := validType.FieldByName(e.Field())
			//反射获取到表单字段
			formField := field.Tag.Get("form")

			//获取分区配置文件
			langText := v.cfg.Section(name).Key(formField).String()

			tranText := e.Translate(v.trans)
			//如果翻译有内容，则替换
			if len(langText) == 0 {
				//搜索全局配置文件
				langText = v.cfg.Section("").Key(formField).String()
			}
			if len(langText) > 0 {
				tranText = strings.Replace(tranText, e.Field(), langText, -1)
			}
			v.Errors[formField] = tranText

			if len(v.Error) == 0 {
				v.Error = tranText
			}
		}
	}
	return v
}

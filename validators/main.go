//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:47

package validators

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/ini.v1"
	"log"
	"reflect"
	"strings"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)


func New() {
	var err error
	chinese := zh.New()
	uni := ut.New(chinese)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	err = zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Validated(valid interface{}) map[string]string {
	if err := validate.Struct(valid); err != nil {
		 return  NewValidatorError(valid,err, trans)
	}
	return  nil
}

func NewValidatorError(valid interface{},err error,trans ut.Translator)  map[string]string {

	errs := err.(validator.ValidationErrors)
	errMap:=make(map[string]string)
	validType := reflect.TypeOf(valid)
	name:=validType.Name()
	for _, e := range errs {
		field, _ := validType.FieldByName(e.Field())
		//反射获取到表单字段
		formField :=field.Tag.Get("form")
		//读取语言配置文件
		langFile:="./resources/langs/zh_CN.ini"
		lang, err := ini.Load(langFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		//获取分区配置文件
		langText := lang.Section(name).Key(formField).String()
		tranText := e.Translate(trans)
		//如果翻译有内容，则替换
		if len(langText) == 0  {
			//搜索全局配置文件
			langText = lang.Section("").Key(formField).String()
		}
		if len(langText) >0 {
			tranText = strings.Replace(tranText, e.Field(), langText, -1)
		}
		errMap[formField] = tranText
	}
	return errMap
}
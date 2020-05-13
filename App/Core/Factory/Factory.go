//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:58

package Factory

import (
	"github.com/gin-gonic/gin"
	"go-gin-project/App/Config/Errors"
	"go-gin-project/App/Core/Container"
	"go-gin-project/App/Http/Request/Interface"
	"log"
	"reflect"
)

//模块获取
func Get(key string)  interface{} {
	if value := Container.CreateContainersFactory().Get(key); value != nil {
		return value
	}
	log.Panicln(Errors.Errors_Container_Not_Exists + ", 模块：" + key)
	return nil
}

// 表单参数工厂
func Create(key string) func(context *gin.Context) {
	if value := Container.CreateContainersFactory().Get(key); value != nil {
		valueOf := reflect.ValueOf(value)
		valueOfInterface := valueOf.Interface()
		//校验是否有CheckParams方法
		if value, ok := valueOfInterface.(Interface.ValidatorInterface); ok {
			return value.CheckParams
		}
	}
	log.Panicln(Errors.Errors_Valiadator_Not_Exists + ", 验证器模块：" + key)
	return nil
}

//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 13:55

package Factory

import (
	"go-gin-project/App/Core/Container"
	"go-gin-project/App/Global/Errors"
	"log"
)

func Get(key string) interface{} {

	if value := Container.CreateContainersFactory().Get(key); value != nil {
		 return value
	}
	log.Panicln(Errors.Errors_Container_Not_Exists + ", 模块：" + key)
	return nil
}
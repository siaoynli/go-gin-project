//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:21

package Bootstrap

import (
	"go-gin-project/App/Core"
	"go-gin-project/App/Global/Errors"
	"go-gin-project/App/Global/Variable"
	"go-gin-project/App/Http/Request/RegisterValidator"
	"log"
	"os"
)

func init(){
	if path, err := os.Getwd(); err == nil {
		Variable.BASE_PATH = path
	} else {
		log.Fatal(Errors.Errors_BasePath)
	}
	Core.InitContainer()
	RegisterValidator.RegisterValidator()
}
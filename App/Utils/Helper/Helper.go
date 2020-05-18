//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 9:30

package Helper

import (
	"fmt"
	"go-gin-project/App/Global/Variable"
	"gopkg.in/ini.v1"
	"log"
)

//加载语言文件
func LoadLang() *ini.File {
	langFile := Variable.BASE_PATH+"/Resources/langs/zh_CN.ini"
	cfg, err := ini.Load(langFile)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("load...")
	return cfg
}



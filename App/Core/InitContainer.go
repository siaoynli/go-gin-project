//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 9:26

package Core

import (
	"go-gin-project/App/Config/Consts"
	"go-gin-project/App/Core/Container"
	"go-gin-project/App/Http/Request/Users"
	"go-gin-project/App/Utils/Helper"
)

// 各个业务模块必须进行注册（初始化），程序启动时会自动加载到容器
func InitContainer(){
	var key string
	//创建容器
	container := Container.CreateContainersFactory()

	//读取语言文件
	container.Set(Consts.Validator_Prefix+"Lang",Helper.LoadLang())

	//用户登录
	key = Consts.Validator_Prefix + "UsersLogin"
	container.Set(key, &Users.Login{})

}
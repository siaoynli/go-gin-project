//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:45

package Kernel

import (
	"go-gin-project/App/Config/Consts"
	"go-gin-project/App/Core/Container"
	"go-gin-project/App/Http/Request"
	"go-gin-project/App/Http/Request/Users"
)

// 各个业务模块必须进行注册（初始化），程序启动时会自动加载到容器
func RegisterKernel(){
	var key string
	//创建容器
	container := Container.CreateContainersFactory()
	//注册多语言
	container.Set(Consts.Validator_Prefix+"Validator",Request.NewValidator())

	key = Consts.Validator_Prefix + "UsersLogin"
	container.Set(key, &Users.Login{})

}
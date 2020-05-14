//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 10:59

package RegisterValidator

import (
	"go-gin-project/App/Core/Container"
	"go-gin-project/App/Global/Consts"
	"go-gin-project/App/Http/Request/Users"
)

func RegisterValidator(){
	var key string
	//创建容器
	container := Container.CreateContainersFactory()
	//用户登录
	key = Consts.Validator_Prefix + "UsersLogin"
	container.Set(key, &Users.Login{})

}
//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:23

package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-project/App/Core/Factory"
	"go-gin-project/App/Global/Consts"
	Users2 "go-gin-project/App/Http/Request/Users"
	"go-gin-project/App/Services"
	"go-gin-project/App/Utils/Response"
	"net/http"
)

type Users struct {
}

func (u *Users) Login(context *gin.Context){

	login := Factory.Get(Consts.Validator_Prefix + "UsersLogin")

	service := Services.NewUserService()

	if service.Store(login.User, login.Password, "lee", "13516872342", "demo") {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, "")
	}else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Creat_Fail_Code, Consts.Curd_Creat_Fail_Msg, "")
	}


}

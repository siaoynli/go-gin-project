//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:26

package Users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-project/App/Config/Consts"
	"go-gin-project/App/Core/Factory"
	"go-gin-project/App/Http/Controllers"
	"go-gin-project/App/Http/Request"
	"go-gin-project/App/Utils/Response"
	"net/http"
)

type Login struct {
	User     string `form:"user" validate:"required,min=5,max=20"`
	Password string `form:"password" validate:"required,min=5,max=20"`
}

func (l Login) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&l); err != nil {
		errs := gin.H{
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}
	Validator := Factory.Get(Consts.Validator_Prefix + "Validator")
	if  err :=Validator.(*Request.Validator).Validated(l).Error;err!= "" {
		fmt.Println("err:",err)
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, err)
	}
	//注册控制器
	(&Controllers.Users{}).Login(context)
}
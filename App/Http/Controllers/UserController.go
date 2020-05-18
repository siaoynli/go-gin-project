//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:23

package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-project/App/Global/Consts"
	"go-gin-project/App/Models/UserModel"
	"strconv"

	"go-gin-project/App/Services"
	"go-gin-project/App/Utils/Response"
	"net/http"
)

type Users struct {
}

func (u *Users) All(context *gin.Context) {
	service := Services.NewUserService()
	Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, gin.H{"lists":service.All()})
}

func (u *Users) Index(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, gin.H{"error": err.Error()})
		return
	}
	service := Services.NewUserService()
	user := service.First(uint(id))
	fmt.Println("user:",user)
	if user.ID == 0 {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, gin.H{"error": "数据不存在"})
		return
	}
	Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, user)
}

func (u *Users) Create(context *gin.Context) {

	login, err := (&UserModel.CreateUsers{}).Validator(context)
	if err != nil {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, gin.H{"error": err.Error()})
		return
	}
	service := Services.NewUserService()
	if service.Create(login) {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, "")
	} else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Creat_Fail_Code, Consts.Curd_Creat_Fail_Msg, "")
	}

}

func (u *Users) Delete(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, gin.H{"error": err.Error()})
		return
	}

	service := Services.NewUserService()
	user := service.First(uint(id))
	if user.ID == 0 {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, gin.H{"error": "数据不存在"})
		return
	}

	service.Delete()
	Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, "")
}

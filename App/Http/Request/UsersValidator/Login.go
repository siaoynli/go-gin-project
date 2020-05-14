//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:26

package UsersValidator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gin-project/App/Http/Request"
)


type LoginForm struct {
	User     string `form:"user" validate:"required,min=5,max=20"`
	Password string `form:"password" validate:"required,min=5,max=20"`
}

func NewLoginForm() *LoginForm{
    return &LoginForm{}
}

func (f *LoginForm) Validator(context *gin.Context) (*LoginForm,error) {
	if err := context.ShouldBind(f); err != nil {
		return &LoginForm{},err
	}
	Validator :=Request.NewValidator()
	if  err :=Validator.Validated(f).Error;err!= "" {
		return &LoginForm{},errors.New(err)
	}
	return f,nil
}

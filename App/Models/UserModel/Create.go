//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 15:01

package UserModel

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-gin-project/App/Http/Request"
	"go-gin-project/App/Models"
)



type CreateUsers struct {
	gorm.Model
	Username     string `form:"username" validate:"required,min=5,max=20"`
	Password string `form:"password" validate:"required,min=5,max=20"`
	Phone    string `form:"phone"`
	RealName string `form:"real_name"`
	Status   int    `form:"status"`
	Token    string
}

func (u *CreateUsers) TableName() string {
	return "users"
}


func (u *CreateUsers) Validator(context *gin.Context) (*CreateUsers,error) {
	if err := context.ShouldBind(u); err != nil {
		return &CreateUsers{},err
	}
	Validator :=Request.NewValidator()
	if  err :=Validator.Validated(u).Error;err!= "" {
		return &CreateUsers{},errors.New(err)
	}
	return u,nil
}


func (u *CreateUsers) Create() bool {
	err := Models.DB.Create(u).Error
	if err != nil {
		return false
	}
	return true
}
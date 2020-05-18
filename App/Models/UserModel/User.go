//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 11:20

package UserModel

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-gin-project/App/Models"
)



type UserModel struct {
	gorm.Model
	Username     string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
	Phone    string `json:"phone"`
	RealName string `json:"real_name"`
	Status   int    `json:"status"`
	Token    string
}

func NewUserModel() *UserModel {
	 return  &UserModel{}
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) All() []*UserModel {
	var users  []*UserModel
	Models.DB.Order("created_at desc").Find(&users)
	return users
}


func (u *UserModel) First(id uint) *UserModel {
	Models.DB.First(u, id)
	return u
}

func (u *UserModel) Delete() bool {
	Models.DB.Delete(u)
	return true
}


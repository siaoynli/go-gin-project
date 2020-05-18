//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 12:06

package Services

import (
	"go-gin-project/App/Models/UserModel"
	"go-gin-project/App/Utils/MD5Cryt"
)

type UserService struct {
	model *UserModel.UserModel
}

func NewUserService() *UserService {
	user:=UserModel.NewUserModel()
	return &UserService{
		model:user,
	}
}

func (u *UserService) All() []*UserModel.UserModel  {
	return u.model.All()
}



func (u *UserService) Create(model *UserModel.CreateUsers) bool {
	// 预先处理密码加密，然后存储在数据库
	//一些逻辑
	model.Password = MD5Cryt.Base64Md5(model.Password)
	model.Status=1
	return model.Create()
}

func (u *UserService) First(id uint)  *UserModel.UserModel {

	model:=UserModel.NewUserModel()
	u.model= model.First(id)
	return  u.model
}

func (u *UserService) Delete()  {
     u.model.Delete()
}
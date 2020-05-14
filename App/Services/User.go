//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 12:06

package Services

import (
	"go-gin-project/App/Models"
	"go-gin-project/App/Utils/MD5Cryt"
)

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Store(name string, pass string, real_name string, phone string, remark string) bool {
	pass = MD5Cryt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return Models.NewUserModel().Store(name, pass, real_name, phone, remark)
}
//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:26

package validators

type LoginForm struct {
	User     string `form:"user" validate:"required,min=5,max=20"`
	Password string `form:"password" validate:"required,min=5,max=20"`
}
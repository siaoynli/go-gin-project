//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:25

package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-project/validators"
)


func Login(c *gin.Context) {
	var loginForm validators.LoginForm

	v := validators.New()
	//绑定错误时候
	_, err := v.Bind(c, loginForm)
	if err != nil {
		c.JSON(500, gin.H{"status": "服务器内部错误"})
	}



	//后面登录逻辑
	if loginForm.User == "user" && loginForm.Password == "password" {
		c.JSON(200, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(401, gin.H{"status": "unauthorized"})
	}
}

//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:25

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-project/validators"

)


func Login(c *gin.Context) {
	var loginForm validators.LoginForm

	//绑定错误时候
	if err:=c.ShouldBind(&loginForm);err != nil {
		fmt.Println(err.Error())
		c.JSON(500,gin.H{
			"status":"http server error",
			"error":err.Error(),
		})
		return
	}

	validators.New()
	validated := validators.Validated(loginForm)
	if  validated!= nil {
		c.JSON(401, gin.H{"message":validated})
	}
	//后面登录逻辑
	if loginForm.User == "user" && loginForm.Password == "password" {
		c.JSON(200, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(401, gin.H{"status": "unauthorized"})
	}
}

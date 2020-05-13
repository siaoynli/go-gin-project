//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:19

package Routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-project/App/Config/Consts"
	"go-gin-project/App/Core/Factory"
	"net/http"
)

func InitRouter() *gin.Engine{
	gin.DisableConsoleColor()
	router := gin.Default()
	//注册中间件
	//处理静态资源（不建议gin框架处理静态资源，参见 Public/readme.md 说明 ）
	router.Static("/public", "./Public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFS("/dir", http.Dir("./Public"))    // 将Public目录内的文件列举展示
	router.StaticFile("/readme", "./Public/readme.md") // 可以根据文件名绑定需要返回的文件名

	router.POST("/login", Factory.Create(Consts.Validator_Prefix+"UsersLogin"))

	return router
}



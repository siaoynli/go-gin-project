//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:19

package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-project/controllers"
)

func RegisterRouter(r *gin.Engine){
	r.GET("/",controllers.Index)
	r.POST("/login",controllers.Login)
}



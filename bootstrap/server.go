//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:21

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-gin-project/router"
)


func Run(){
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run(":8080")
}
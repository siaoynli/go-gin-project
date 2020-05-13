//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:23

package Controllers

import (
	"github.com/gin-gonic/gin"


)

type Users struct {
}

func (u *Users) Login(c *gin.Context){
	c.JSON(200, gin.H{"status": "ok"})
}

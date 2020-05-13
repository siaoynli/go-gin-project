//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:17

package controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.String(200,"hello world")
}

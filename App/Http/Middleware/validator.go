//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 14:42

package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-project/utils"
)

func ValidatorMiddleware(c *gin.Context) {

	validator := utils.NewValidator()
	c.Set("validator",validator)
	c.Next()
}


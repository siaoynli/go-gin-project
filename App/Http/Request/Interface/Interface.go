//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 17:00

package Interface

import "github.com/gin-gonic/gin"

type ValidatorInterface interface {
	Validator(context *gin.Context)
}



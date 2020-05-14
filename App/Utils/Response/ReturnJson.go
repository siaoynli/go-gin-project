package Response

import (
	"github.com/gin-gonic/gin"
)

func ReturnJson(context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	// 返回 json数据
	context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

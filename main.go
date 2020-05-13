//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 9:13

package main

import (
     _ "go-gin-project/Bootstrap"
     "go-gin-project/Routers"
)

func main(){
    routers := Routers.InitRouter()
    routers.Run(":8080")
}



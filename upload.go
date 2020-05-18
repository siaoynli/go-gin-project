//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/15 14:25

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"io"
	"os"
	"time"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
}

func main() {
	r := gin.Default()
	InitRedis()
	r.POST("/upload", func(context *gin.Context) {
		total := 0
		file, header, err := context.Request.FormFile("file")
		if err != nil {
			fmt.Println(err)
			context.String(200, "上传文件有误")
			return
		}
		saveFile, err := os.OpenFile(header.Filename, os.O_RDWR, os.ModePerm)
		if err != nil {
			saveFile, _ = os.Create(header.Filename)
		}
		buf := make([]byte, 2048)
		//获取先前上传过的位置记录
		if result, err := Client.Get(header.Filename).Int64(); err == nil {
			fmt.Println("先前已经上传过:", result)
			total = int(result)
			//先前上传过的位置，从文件开头位置偏移total
			file.Seek(int64(total), io.SeekStart)
		}
		num := 0
		for {
			readNum, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			fmt.Println("正在上传，字节数:", num, readNum)
			//跳过total字节追加
			saveFile.WriteAt(buf, int64(total))
			total += readNum
			num += 1
			err = Client.Set(header.Filename, total, time.Duration(0)).Err()
			if err != nil {
				fmt.Println("redis 保存字节数失败")
			}
		}

		fmt.Println("上传完毕", total)
		saveFile.Close()
	})

	r.Run(":8080")
}

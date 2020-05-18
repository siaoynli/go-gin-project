//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 14:26

package Models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-project/App/Global/Errors"
	"go-gin-project/App/Utils/Config"
	"log"
)

var DB *gorm.DB

func init() {
	if DB == nil {
		DB=InitConn()
	}
}

func InitConn() *gorm.DB{
	configFac := Config.CreateYamlFactory()
	DbType := configFac.GetString("DbType")
	Host := configFac.GetString("Mysql.Host")
	Port := configFac.GetString("Mysql.Port")
	User := configFac.GetString("Mysql.User")
	Pass := configFac.GetString("Mysql.Pass")
	Charset := configFac.GetString("Mysql.Charset")
	DataBase := configFac.GetString("Mysql.DataBase")

	db, err := gorm.Open(DbType, (User)+":"+Pass+"@tcp("+Host+":"+Port+")/"+DataBase+"?parseTime=True&loc=Local&charset="+Charset)
	if err != nil {
		log.Fatal(Errors.Errors_Db_SqlDriverInitFail)
	}
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	//db.SingularTable(true)
	//这里关闭了，会造成无法访问数据，所以设置连接池即可
	//defer db.Close()
	db.DB().SetMaxIdleConns(configFac.GetInt("Mysql.SetMaxIdleConns"))
	db.DB().SetMaxOpenConns(configFac.GetInt("Mysql.SetMaxOpenConns"))
	db.DB().SetConnMaxLifetime(configFac.GetDuration("Mysql.SetConnMaxLifetime"))
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// 自动迁移模式
	//db.AutoMigrate(&UserModel.UsersModel{})
	fmt.Println("DB load...")
	return db
}
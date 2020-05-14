//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 12:02

package Factory

import (
	"database/sql"
	"go-gin-project/App/Global/Errors"
	"go-gin-project/App/Utils/Config"
	"log"
	"time"
)

// 初始化数据库驱动
func InitSqlDriver() *sql.DB {
	configFac := Config.CreateYamlFactory()
	DbType := configFac.GetString("DbType")
	Host := configFac.GetString("Mysql.Host")
	Port := configFac.GetString("Mysql.Port")
	User := configFac.GetString("Mysql.User")
	Pass := configFac.GetString("Mysql.Pass")
	Charset := configFac.GetString("Mysql.Charset")
	DataBase := configFac.GetString("Mysql.DataBase")
	SetMaxIdleConns := configFac.GetInt("Mysql.SetMaxIdleConns")
	SetMaxOpenConns := configFac.GetInt("Mysql.SetMaxOpenConns")
	SetConnMaxLifetime := configFac.GetDuration("Mysql.SetConnMaxLifetime")
	db, err := sql.Open(DbType, string(User)+":"+Pass+"@tcp("+Host+":"+Port+")/"+DataBase+"?parseTime=True&loc=Local&charset="+Charset)
	if err != nil {
		log.Fatal(Errors.Errors_Db_SqlDriverInitFail)
	}
	db.SetMaxIdleConns(SetMaxIdleConns)
	db.SetMaxOpenConns(SetMaxOpenConns)
	db.SetConnMaxLifetime(SetConnMaxLifetime * time.Second)

	// 将需要销毁的事件统一注册在事件管理器，由程序退出时统一销毁
	//Event.CreateEventManageFactory().Set(Variable.Event_Destroy_Prefix+"DB", func(args ...interface{}) {
	//	db.Close()
	//})
	return db
}

// 从连接池获取一个连接
func GetOneEffectivePing() *sql.DB {
	configFac := Config.CreateYamlFactory()
	maxRetryTimes := configFac.GetInt("Mysql.PingFailRetryTimes")
	// ping 失败允许重试
	vDbDriver := InitSqlDriver()
	for i := 1; i <= maxRetryTimes; i++ {
		if err := vDbDriver.Ping(); err != nil { //  获取一个连接失败，进行重试
			vDbDriver = InitSqlDriver()
			time.Sleep(time.Second * 1)
			if i == maxRetryTimes {
				log.Fatal(Errors.Errors_Db_GetConnFail)
			}
		} else {
			break
		}
	}
	return vDbDriver
}


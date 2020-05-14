//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 12:00

package Models

import (

	"database/sql"
	"go-gin-project/App/Core/Factory"
	"go-gin-project/App/Global/Errors"
	"log"
)

var sqlDriver *sql.DB

// 创建一个数据库基类工厂
func CreateBaseSqlFactory(sqlMode string) (res *BaseModel) {
	switch sqlMode {
	case "mysql":
		if sqlDriver == nil {
			sqlDriver = Factory.InitSqlDriver()
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := sqlDriver.Ping(); err != nil {
			// 重试
			sqlDriver = Factory.GetOneEffectivePing()
			// 如果重试成功
			if err := sqlDriver.Ping(); err == nil {
				res = &BaseModel{dbDriver: sqlDriver}
			}
		} else {
			res = &BaseModel{dbDriver: sqlDriver}
		}
	default:
		log.Panic(Errors.Errors_Db_Driver_NotExists)
	}
	return res
}

// 定义一个数据库操作的基本结构体
type BaseModel struct {
	dbDriver *sql.DB
}

//  执行类: 新增、更新、删除
func (b *BaseModel) ExecuteSql(sql string, args ...interface{}) int64 {
	if stm, err := b.dbDriver.Prepare(sql); err == nil {
		if res, err := stm.Exec(args...); err == nil {
			if affectNum, err := res.RowsAffected(); err == nil {
				return affectNum
			} else {
				log.Panic("执行类sql失败:", err)
			}
		} else {
			log.Panic("(预处理)执行类sql失败:", err)
		}
	}
	return -1

}

//  查询类: select
func (b *BaseModel) QuerySql(sql string, args ...interface{}) *sql.Rows {
	if stm, err := b.dbDriver.Prepare(sql); err == nil {
		// 可变参数的二次传递，需要在后面添加三个点 ...  ，这里和php刚好相反
		if Rows, err := stm.Query(args...); err == nil {
			return Rows
		}
	}
	return nil

}



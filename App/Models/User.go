//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 11:20

package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type UsersModel struct {
	*BaseModel
	Id       int64  `json:"id"`
	User      string `json:"username"`
	Password     string `json:"pass"`
	Phone    string `json:"phone"`
	RealName string `json:"real_name"`
	Status   int    `json:"status"`
	Token    string
}

func NewUserModel() *UsersModel {
	dbDriver := CreateBaseSqlFactory("mysql")
	if dbDriver != nil {
		return &UsersModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("UsersModel工厂初始化失败")
	return nil
}


func (u *UsersModel) Store(username string, pass string, realName string, phone string, remark string) bool {
	sql := "INSERT  INTO tb_users(username,pass,real_name,phone,remark) SELECT ?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  username=?)"
	if u.ExecuteSql(sql, username, pass, realName, phone, remark, username) > 0 {
		return true
	}
	return false
}



//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:38

package Errors

const (
	//系统部分
	Errors_FuncEvent_Already_Exists string = "注册函数类事件失败，键名已经被注册"
	Errors_FuncEvent_NotRegister    string = "没有找到键名对应的函数"
	Errors_FuncEvent_NotCall        string = "注册的函数无法正确执行"
	Errors_BasePath                 string = "初始化项目根目录失败"
	Errors_NoAuthorization          string = "token鉴权未通过，请通过token授权接口重新获取token,"
	// 数据库部分
	Errors_Db_Driver_NotExists  string = "数据库驱动不存在"
	Errors_Db_SqlDriverInitFail string = "数据库驱动初始化失败"
	Errors_Db_GetConnFail       string = "从数据库连接池获取一个连接失败，超过最大连接次数."
	//redis部分
	Errors_Redis_InitConnFail string = "初始化redis连接池失败"
	Errors_Redis_AuhtFail     string = "Redis Auht鉴权失败，密码错误"

	Errors_Container_Not_Exists string = "不存在的模块"
	// 验证器错误
	Errors_Valiadator_Not_Exists string = "不存在的验证器"

	//token部分
	Errors_Token_Invalid string = "无效的token"
	Errors_Token_Expired string = "过期的token"
	//
	Errors_Snowflake_Init_Fail  string = "初始化 snowflakeFctory 过程发生错误"
	Errors_Snowflake_GetId_Fail string = "获取snowflake唯一ID过程发生错误"
)




//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:17

package Variable

var (
	// 系统预设全库变量
	BASE_PATH            string              // 定义项目的根目录
	Event_Destroy_Prefix string = "destroy_" //  程序退出时需要销毁的事件前缀
	//上传文件保存路径
	UploadFileField    string = "files"                  // post上传文件时，表单的键名
	UploadFileSavePath string = "/Storage/app/uploaded/" // 该路径与 base_path 进行拼接使用
	//日志存储路径
	Log_Save_Path string = "/Storage/logs/gin.log"

	//  用户自行定义其他全局变量 ↓

)




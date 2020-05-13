//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/13 16:41

package Container

import (
	"strings"
	"sync"
)

// 定义一个全局键值对存储容器
var syncMap sync.Map

// 定义一个容器结构体
type containers struct {
}

// 创建一个容器工厂
func CreateContainersFactory() *containers {
	return &containers{}
}

//  1.以键值对的形式将代码注册到容器
func (e *containers) Set(key string, value interface{}) (res bool) {
	if e.Get(key) == nil {
		syncMap.Store(key, value)
		res = true
	}
	return
}

//  2.删除
func (e *containers) Delete(key string) {
	syncMap.Delete(key)
}

//  3.传递键，从容器获取值
func (e *containers) Get(key string) interface{} {

	if value, exists := syncMap.Load(key); exists {
		return value
	}
	return nil
}

// 按照键的前缀模糊删除容器中注册的内容
func (e *containers) FuzzyDelete(keyPrefix string) {
	syncMap.Range(func(key, value interface{}) bool {
		if keyName, ok := key.(string); ok {
			if strings.HasPrefix(keyName, keyPrefix) {
				syncMap.Delete(keyName)
			}
		}
		return true
	})
}




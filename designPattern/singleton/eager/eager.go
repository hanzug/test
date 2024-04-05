package eager

type singleton struct{}

// 初始化单例实例，由于是包级别的变量，会在包初始化时创建，属于饿汉模式。
var instance = &singleton{}

// GetInstance 提供一个获取单例对象的方法
func GetInstance() *singleton {
	return instance
}

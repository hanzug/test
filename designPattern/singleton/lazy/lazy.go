package lazy

import (
	"sync"
)

type Singleton struct{}

var instance *Singleton

var mu sync.Mutex

// GetInstance 双重检查加锁的懒汉模式
func GetInstanceLazy() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}

//	 sync.Once更加简洁

//var instance *singleton
//var once sync.Once
//
//// GetInstance 使用sync.Once的懒汉模式
//func GetInstance() *singleton {
//	once.Do(func() {
//		instance = &singleton{}
//	})
//	return instance
//}

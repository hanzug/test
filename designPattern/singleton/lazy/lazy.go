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
		//这里会出现两个线程同时进入的情况。然后他们就会竞争锁。
		mu.Lock()
		defer mu.Unlock()
		//得到锁的线程再次进行判断，防止多次创建。
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

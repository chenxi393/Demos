package main

import "sync"

// 单例（Singleton）模式
var svcOnce sync.Once
var singleton *LazySingleton

// 懒汉式(等到调用的时候才封装，使用golang现成的once方法防止并发问题)
func GetSingletonByLazy() *LazySingleton {
	svcOnce.Do(func() {
		singleton = &LazySingleton{}
	})
	return singleton
}

type LazySingleton struct {
}

// 饿汉式(代码加载的时候就封装好单例而非等到调用的时候才封装)
func init() {
	singleton = &LazySingleton{}
}

func GetSingleton() *LazySingleton {
	return singleton
}

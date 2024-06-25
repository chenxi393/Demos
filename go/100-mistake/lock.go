package main

import (
	"sync"
	"time"
)

func main() {
	counter := NewCounter()
 
	go func() {
	   counter.Increment1("foo")
	}()
	go func() {
	   counter.Increment1("bar")
	}()
 
	time.Sleep(10 * time.Millisecond)
 }
 
 type Counter struct {
	mu       sync.Mutex
	counters map[string]int
 }
 
 func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
 }
 
 // 复制sync.Mutex 加的不是同一个锁
 // 实际上 因为是值类型 一个函数Increment1内的修改 并不会 影响另一个 Increment1 的调用
 // 也就是 锁的状态是不会变的 
 func (c *Counter) Increment1(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++ // 并发读写错误
 }
 //  go run -race xx.go 存在数据竞争
package pool

import (
	"fmt"
	"sync"
	"time"
)

// 协程池 的实现 其实和线程池一样
// 可以考虑 不使用queue 直接使用
type Work struct {
	f   func(interface{})
	arg interface{}
}

// 是不是应该基于 信号量的 唤醒机制
// 看下 go pool ants 的实现
type Poll struct {
	queue []Work
	ch    chan bool
	lock  sync.Mutex
}

func (p *Poll) InitPoll(num int) {
	// 似乎channel 指定缓存不是一个好的习惯
	p.ch = make(chan bool, num*num)
	p.queue = make([]Work, 0, num*num)
	for i := 0; i < num; i++ {
		go func() {
			for {
				hasWork, ok := <-p.ch
				if !ok {
					break
				}
				if hasWork {
					p.lock.Lock()
					w := p.queue[0]
					p.queue = p.queue[1:]
					p.lock.Unlock()
					w.f(w.arg)
				}
			}
		}()
	}
}

func (p *Poll) AddWork(w Work) {
	p.lock.Lock()
	p.queue = append(p.queue, w)
	p.lock.Unlock()
	p.ch <- true // 这里发送会阻塞 其实要让work 自己探究 有没有阻塞的能力 可以探究一下 go_pool 的实现
}

func Test_Poll() {
	var p Poll
	p.InitPoll(1000)
	wg := sync.WaitGroup{}
	w := &Work{
		f: func(i interface{}) {
			time.Sleep(1 * time.Second)
			fmt.Println("Now print:", i)
			wg.Done()
		},
	}
	wg.Add(800)
	for i := 0; i < 800; i++ {
		w.arg = i
		p.AddWork(*w)
	}
	wg.Wait()
	// var i int
	// for {
	// 	fmt.Scan(&i)
	// 	w.arg = i
	// 	p.AddWork(*w)
	// }
}

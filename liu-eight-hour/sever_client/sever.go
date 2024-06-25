package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Sever struct {
	Ip   string
	Port string

	//2.0新增
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

func InitSever(ip string, port string) *Sever {
	s := &Sever{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return s
}

func (s *Sever) ListenSever() {
	for {
		msg := <-s.Message

		//发送给每一个
		s.mapLock.Lock()
		for _, cil := range s.OnlineMap {
			cil.C <- msg
		}
		s.mapLock.Unlock()
	}
}

func (s *Sever) Brordcast(u *User, msg string) {
	message := "[" + u.Addr + "]  :" + u.Name + " " + msg
	s.Message <- message
}

func (s *Sever) Handle(c net.Conn) {
	//当前连接的业务
	fmt.Println(c.RemoteAddr().String() + " is connected.")

	//监听用户活跃的channel
	islive := make(chan bool)

	//先创建用户对象
	user := InitUser(c, s)

	//加入到在线用户表里 要加读写锁
	user.Online()

	//这里是否需要阻塞呢
	//视频里用select阻塞了
	//不阻塞的化 user是否会被回收 里面的go程是否运行

	//3.0 新增  接受客户端发送的消息
	go func() {
		buff := make([]byte, 4096)
		for {
			n, err := c.Read(buff)
			//下线的时候会发长度为0的消息
			if n == 0 {
				//这样的写法会造成服务端出现大量CLOSE_WAIT，return也只是退出当前协程，而不是Handle
				user.Offline()
				return
			}
			//读到末尾的时候err是io.EOF
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read Error: ", err.Error())
				return
			}

			msg := string(buff[:n-1])

			user.Message_(msg)

			islive <- true
		}
	}()
	for {
		select {
		case <-islive:
			//当前用户是活跃的,不做任何事情，select中会执行下面这句重置
			//time.After(time.Second * 10)重新执行即刻重置定时器，定时到后会发送信息
			//这里select 第二个case定时器触发后，处于阻塞状态。当满足第一个 case 的条件后，
			//打破了 select 的阻塞状态，每个条件又开始判断，第2个 case 的判断条件一执行，就重置定时器了。
		case <-time.After(time.Minute):
			user.send_message("你被强制下线\n")

			//销毁资源//只关闭uer.C
			close(user.C)
			return
		}
	}

}

func (s *Sever) SeverStart() {
	//socked listen
	listen, err := net.Listen("tcp", s.Ip+":"+s.Port) //fmt.Sprintf() 用这个拼接也可以
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	//close the socket
	defer listen.Close()

	go s.ListenSever()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err: ", err)
			continue
		}

		go s.Handle(conn)
		//这个用go 不用函数 因为后面会创建子go程 这里用函数循环结束 后面goroutine全挂
		//似乎上面这句是错的 实测没有go也可以一直运行 协程无父子关系 父进程退出不影响子进程
		//只要main不挂就行  加go更多的优化性能吧 可以并发
	}
}

package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	SS *Sever
}

func (u *User) UserListen() { //这种写法是当前user的方法 无参数 无返回值
	//5.0 修复bug 当u.C关闭后 不再接受消息 退出for
	for msg := range u.C {
		_, err := u.conn.Write([]byte(msg + "\n"))
		if err != nil {
			panic(err)
		}
	}

	err := u.conn.Close()
	if err != nil {
		panic(err)
	}
}

func (u *User) Online() {
	u.SS.mapLock.Lock()
	u.SS.OnlineMap[u.Name] = u
	u.SS.mapLock.Unlock()

	u.SS.Brordcast(u, "已上线")
}

func (u *User) Offline() {
	u.SS.mapLock.Lock()
	delete(u.SS.OnlineMap, u.Name)
	u.SS.mapLock.Unlock()

	u.SS.Brordcast(u, "已下线")
}

func (u *User) Message_(msg string) {
	//4.0 查询当前
	if msg == "check" {
		u.check_all_users()
	} else if strings.Contains(msg, "rename:") {
		u.rename(msg[7:])
		//msg 也可以 strings.Split(msg,":")[1]
	} else if len(msg) >= 3 && msg[0:3] == "TO:" { //新增私聊功能
		pri_name := strings.Split(msg, ":")[1]
		pri_msg := strings.Split(msg, ":")[2]
		if pri_name == "" {
			u.send_message("请输入用户名\n")
			return
		}
		//在if语句块外部，变量ok是不可见的
		//这样叫做短路赋值 但是可读性不好 
		if to_user, ok := u.SS.OnlineMap[pri_name]; ok {
			to_user.send_message(u.Name+"对你说："+pri_msg+"\n")
		} else {
			u.send_message("用户名不存在\n")
		}
	} else {
		u.SS.Brordcast(u, msg)
	}
}

// 给自己的端口写消息
func (u *User) send_message(msg string) {
	u.conn.Write([]byte(msg))
}

// 4.0 新增查询当前用户
func (u *User) check_all_users() {
	u.SS.mapLock.RLock()
	i := 1
	for _, add := range u.SS.OnlineMap {
		u.send_message(fmt.Sprint(i) + ": " + add.Name + " is online.\n")
		i++
	}
	u.SS.mapLock.RUnlock()
}

// 5.0 新增改名
func (u *User) rename(New string) {
	u.SS.mapLock.Lock()
	//先看当前New 在不在map里
	_, ok := u.SS.OnlineMap[New]
	if ok {
		u.send_message("更改失败：用户名已存在\n")
	} else {
		delete(u.SS.OnlineMap, u.Name)
		u.SS.OnlineMap[New] = u
		u.send_message("更改成功\n")
		u.Name = New
	}
	u.SS.mapLock.Unlock()
}

func InitUser(c net.Conn, S *Sever) (u *User) {
	u = &User{
		Name: c.RemoteAddr().String(),
		Addr: c.RemoteAddr().String(),
		C:    make(chan string),
		conn: c,
		SS:   S,
	}

	go u.UserListen()

	return //这里其实可以证明 协程无父子关系 return了 上面的goroutine 也不会挂
}

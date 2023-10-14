package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type client struct {
	IpAddess string
	Port     int
	usr_name string
	conn     net.Conn
	flag     int //模式的选择
}

var global_ip string
var global_port int

// ./client -ip 127.0.0.1 -port 8880
// 解析命令行
func init() {
	flag.StringVar(&global_ip, "ip", "127.0.0.1", "设置服务器IP")
	flag.IntVar(&global_port, "port", 8888, "设置IP的端口")
}

func New_client(ip string, port int) *client {
	//1.创建客户端对象
	usr := &client{
		IpAddess: ip,
		Port:     port,
		usr_name: "defalt",
		flag:     999,
	}

	//连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))

	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	usr.conn = conn
	//返回对象
	return usr
}

func (usr *client) DealResponse() {
	// for {
	// 	var msg = make([]byte, 100)
	// 	usr.conn.Read(msg)
	// 	fmt.Println(msg)
	// }

	//这里等价于上面
	io.Copy(os.Stdout, usr.conn)
	//永久阻塞的 conn一旦有数据就拷贝到stdout

}

func (usr *client) Rename() {
	fmt.Println(">>>>>请输入你要更改的用户名：")
	fmt.Scanln(&usr.usr_name)

	_, err := usr.conn.Write([]byte("rename:" + usr.usr_name + "\n"))
	if err != nil {
		fmt.Println("conn Writer err:", err)
		return
	}
	//需要接受服务端的消息 得知是否更改成功
	//应当实现并发操作 不要阻塞 使用协程解决
}

func (usr *client) Public_Chat() {
	msg := ""

	for {
		fmt.Println(">>>>>请输入你要发送的消息: exit表示退出")

		fmt.Scanln(&msg)
		if msg == "exit" {
			return
		}
		_, err := usr.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("conn Writer err:", err)
			return
		}
	}

}

func (usr *client) Check_Users() {
	_, err := usr.conn.Write([]byte("check\n"))

	if err != nil {
		fmt.Println("conn Writer err:", err)
		return
	}
}

func (usr *client) Private_Chat() {
	//先给客户端所有的用户在线列表 也就是check功能 实际上也可以做到菜单里

	fmt.Println("在线用户列表")
	usr.Check_Users()

	var name, msg string
	for {
		fmt.Println(">>>>>请输入你要私聊的用户名: exit退出")
		fmt.Scanln(&name)
		if name == "exit" {
			break
		}
		for {
			fmt.Println(">>>>>请输入你要私聊的内容: exit退出")
			fmt.Scan(&msg) //要注意读取空格的问题
			if msg == "exit" {
				break
			} else if msg == "" {
				fmt.Println("不可以发空消息")
			}
			_, err := usr.conn.Write([]byte("TO:" + name + ":" + msg + "\n"))
			if err != nil {
				fmt.Println("conn Writer err:", err)
				return
			}
			msg = ""
		}
	}

}

func (usr *client) Run() {
	for usr.flag != 0 {
		if !usr.menu() {
			fmt.Println(">>>>>请输入正确的数字")
			continue
		}
		switch usr.flag {
		case 1:
			usr.Rename()
		case 2:
			usr.Public_Chat()
		case 3:
			usr.Private_Chat()
		}
		//usr.conn
		//还有一个问题是 要是超时被服务端踢了
		//客户端如何知道
	}
}

func (usr *client) menu() bool {
	fmt.Println(">>>>>请选择你需要的模式")
	fmt.Println("1.更改用户名")
	fmt.Println("2.公聊模式")
	fmt.Println("3.私聊模式")
	fmt.Println("0.退出程序")

	fmt.Scanln(&usr.flag)
	if usr.flag < 0 || usr.flag > 3 {
		return false
	}

	return true
}

func main() {
	//解析命令行参数
	flag.Parse()

	user := New_client(global_ip, global_port)
	if user == nil { //这里是等于nil 说明返回空指针
		fmt.Println(">>>>>客户端连接失败")
		return
	}

	fmt.Println(">>>>>客户端连接成功")
	//启动客户端的业务

	//go单独处理服务器的消息
	go user.DealResponse()

	user.Run()

}

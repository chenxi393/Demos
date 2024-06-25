package main

/*
整体框架
sever.go 服务端
InitSever 程序入口 服务端开启
SeverStart 建立套接字监听
go协程监听channel 给客户端发消息
for循环等待客户端发送连接请求
go s.Handle(conn) 协程处理客户端连接

handle省略说明 主要处理客户端的请求
*/


func main(){
	s:=InitSever("127.0.0.1","8888")
	s.SeverStart()
}
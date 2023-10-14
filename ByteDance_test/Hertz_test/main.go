package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Test() {
	time.Sleep(time.Second * 5)
	c, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	status, body, err := c.Get(context.Background(), nil, "www.baidu.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(status, body, err)

}
func main() {
	h := server.Default(server.WithHostPorts("localhost:1111"))
	// default 会有默认的recover中间价

	// 注意这里有两个上下文 一个用于传递原信息 一个用于请求处理
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.GET("/", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "Hello World!!!")
	})

	go Test()
	// 支持分组路由 参数路由 通配路由
	// 路由优先级 静态路由 > 命名路由 > 通配路由
	//中间件一般可以注册在全局的
	h.Spin()

}

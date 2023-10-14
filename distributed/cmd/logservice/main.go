package main

import (
	"context"
	"dis/log"
	"dis/registry"
	"dis/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000" //正常来说这些东西要放在配置文件里
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start( // 会返回一个子ctx
		context.Background(), //这是父ctx 空白的
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	// 当调用cancel() 函数时  channel会发送信号
	// 具体机制有待了解  TODO
	<-ctx.Done()

	fmt.Println("Shutting down log service.")
}

package service

import (
	"context"
	"dis/registry"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Start(ctx context.Context, host, port string, reg registry.Registration,
	registerHandldersFunc func()) (context.Context, error) {
	registerHandldersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx) // 这个第一次用
	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx) // shutdown会在上面关闭服务
		time.Sleep(time.Second) // 没有这个 会调用cancel直接main退出了
		// 不会等到上面一个gorountine的错误打印
		cancel()
	}()
	return ctx
}

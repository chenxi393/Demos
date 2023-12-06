package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "grpctest/proto"

	eclient "go.etcd.io/etcd/client/v3"
	eresolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// grpc 服务名
	MyService = "douyin/sayhello"
	// etcd 端口
	MyEtcdURL = "http://localhost:2379"
)

func main() {
	// 创建 etcd 客户端
	etcdClient, err := eclient.NewFromURL(MyEtcdURL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// 创建 etcd 实现的 grpc 服务注册发现模块 resolver
	// 然后在调用 grpc.Dial 方法创建连接代理 ClientConn 时，将其注入其中.
	// 类似于一个域名解析器
	etcdResolverBuilder, err := eresolver.NewBuilder(etcdClient)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// 拼接服务名称，需要固定义 etcd:/// 作为前缀
	etcdTarget := fmt.Sprintf("etcd:///%s", MyService)

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		// 服务名称
		etcdTarget,
		// 注入 etcd resolver
		grpc.WithResolvers(etcdResolverBuilder),
		// 声明使用的负载均衡策略为 roundrobin
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSayHelloClient(conn)

	wg := &sync.WaitGroup{}
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func() {
			r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "小明"})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.GetReply())
			wg.Done()
		}()
	}
	wg.Wait()
}

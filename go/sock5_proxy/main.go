package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/proxy"
	// clientv3 "go.etcd.io/etcd/client/v3"
)

func testSock5() {
	// SOCKS5 proxy address
	proxyStr := "172.21.16.1:10808"

	// Create a SOCKS5 dialer
	dialer, err := proxy.SOCKS5("tcp", proxyStr, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 dialer: %v", err)
	}

	// Create an HTTP client using the SOCKS5 dialer
	transport := &http.Transport{
		Dial: dialer.Dial,
	}
	client := &http.Client{
		Transport: transport,
	}

	// Make a request
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Printf("%s\n", body)
}

func testSock() {
	urli := url.URL{}
	urlproxy, _ := urli.Parse("socks://172.21.16.1:10808")
	c := http.Client{
		Transport: &http.Transport{
			// 默认不支持sock5
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	if resp, err := c.Get("https://www.google.com"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}

func main() {
	testSock5()
	return
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"47.115.202.192:2379"},
		DialTimeout: 5 * time.Second,
	})

	// 是代理的锅
	// cli, err := clientv3.New(clientv3.Config{
	// 	Endpoints: []string{"47.115.202.192:2379"},
	// 	// DialOptions: []grpc.DialOption{grpc.WithBlock()},
	// 	DialTimeout: 5 * time.Second,
	// })

	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	_, err = cli.Put(ctx, "q1mi", "dsb")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}

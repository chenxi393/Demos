## 基本流程
[gRPC](https://grpc.io/docs/languages/go/basics/)
1. 编写proto文件 定义好CS通信的service和message
2. 使用`protoc --go_out=dir --go-grpc_out=dir hello.proto` 生成代码
   * hello.pb.go包含用于填充、序列化和检索请求和响应消息类型的所有协议缓冲区代码。
   * hello_grpc_pb.go 服务端要实现的和客户端调用的代码
3. 编写客户端和服务端业务代码 完成rpc调用

## gRPC 认证安全传输
* SSL/TLS认证方式 采用HTTP2
* 基于Token的认知方式（其实就是客户端传递token给服务端验证）
* 不采取任何认证 HTTP1
* 自定义的身份认证

## HTTPS的实现原理
1. HTTPS，SSL，TLS的含义
    SSL和TLS是一个东西的不同阶段，TLS是标准化后的SSL
2. HTTP长/短连接
   * HTTP1.0 需要使用Keep-Alive保持长连接 而HTTP1.1默认就是长连接
3. 加密算法的概念（对称加密和非对称加密）
   * 对称加密：加密和解密的密钥是一样的，性能好点
   * 非对称加密：加密解密是不同的密钥，需要算法生成，更安全，性能差些
4. CA证书的用法
   * 一系列操作生成一堆文件 私钥 证书什么的

## rpc vs http
rpc是基于SDK方式调用，需要调用双方约定好传输协议，通信内容
rpc基于tcp（或者udp），上层协议（应用层）是需要自己去定义的
rpc更使用于内部调用

rpc只能说算是自定义的通信方式，不算传统的网络协议
rpc可以基于tcp，udp，http，甚至别的协议

例如gRPC基于以HTTP2 作为应用层协议

## grpc
protobuf 作为数据序列化协议以及接口定义语言（IDL）
这里引申一下 trift也是一种IDL 接口定义语言

``` sh
protoc --version 
# 基于.proto文件一键生成 _pb.go文件 通信请求/响应参数的对象模型.
protoc-gen-go --version
# 基于 .proto 文件生成 _grpc.pb.go，对应内容为通信服务框架代码.
protoc-gen-go-grpc --version

protoc --go_out=. --go-grpc_out=. pb/hello.proto
# --go_out：指定 pb.go 文件的生成位置
# --go-grpc_out：指定 grpc.pb.go 文件的生成位置
# pb/hello.proto：这是指定了 .proto 文件的所在位置
```

```proto
 rpc ListFeatures(Rectangle) returns (stream Feature) {}
    // Accepts a stream of Points on a route being traversed, returning a
    // RouteSummary when traversal is completed.
    rpc RecordRoute(stream Point) returns (RouteSummary) {}

    // 服务端和客户端均支持流式传输 在数据比较大的时候好用
    // 例如 客户端一点点打印数据 而不用等待服务端全部返回
    // Accepts a stream of RouteNotes sent while a route is being traversed,
    // while receiving other RouteNotes (e.g. from other users).
    rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
```

## grpc 结合 etcd作为服务注册与发现功能
[参考文章 代码已经编写好了](https://etcd.io/docs/v3.5/dev-guide/grpc_naming/)
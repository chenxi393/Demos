## 基本流程
1. 编写proto文件 定义好CS通信的service和message
2. 使用`protoc --go_out=dir --go-rpc_out=dir hello.proto` 生成代码
   * hello.pb.go包含用于填充、序列化和检索请求和响应消息类型的所有协议缓冲区代码。
   * hello_grpc_pb.go 服务端要实现的和客户端调用的代码
3. 编写客户端和服务端业务代码 完成rpc调用

## gRPC 认证安全传输
* SSL/TLS认证方式 采用HTTP2
* 基于Token的认知方式（其实就是客户端传递token给服务端验证）
* 不采取任何认证 HTTP1
* 自定义的身份认证

### HTTPS的实现原理
1. HTTPS，SSL，TLS的含义
    SSL和TLS是一个东西的不同阶段，TLS是标准化后的SSL
2. HTTP长/短连接
   * HTTP1.0 需要使用Keep-Alive保持长连接 而HTTP1.1默认就是长连接
3. 加密算法的概念（对称加密和非对称加密）
   * 对称加密：加密和解密的密钥是一样的，性能好点
   * 非对称加密：加密解密是不同的密钥，需要算法生成，更安全，性能差些
4. CA证书的用法
   * 一系列操作生成一堆文件 私钥 证书什么的

## 官方文档
[gRPC](https://grpc.io/docs/languages/go/basics/)
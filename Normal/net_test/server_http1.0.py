import socket

# 定义服务器地址和端口号
server_address = ('127.0.0.1', 8000)

# 创建一个 TCP socket
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# 绑定服务器地址和端口
server_socket.bind(server_address)

# 开始监听客户端连接
server_socket.listen(1)
print('服务器已启动，监听端口 8000...')

while True:
    # 接受客户端连接
    client_socket, client_address = server_socket.accept()
    
    # 接收客户端的请求数据
    request = client_socket.recv(4096)
    
    # 构造 HTTP 响应
    response = "HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, World!"
    
    # 发送 HTTP 响应给客户端
    client_socket.send(response.encode())
    
    # 关闭客户端连接
    client_socket.close()
import socket

# 定义服务器地址和端口号
server_address = ('localhost', 8000)

# 创建一个 TCP socket
client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# 连接到服务器
client_socket.connect(server_address)

# 构造 HTTP 请求
request = "GET / HTTP/1.0\r\nHost: localhost\r\n\r\n"

# 发送 HTTP 请求给服务器
client_socket.send(request.encode())

# 接收服务器的响应数据
response = client_socket.recv(4096)

# 打印响应数据
print(response.decode())

# 关闭客户端 socket 连接
client_socket.close()
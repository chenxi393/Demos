import socket

UDP_IP = "127.0.0.1"
UDP_PORT = 8888
BUFFER_SIZE = 1024

# 创建 UDP 套接字
server_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

# 绑定 IP 地址和端口号
server_socket.bind((UDP_IP, UDP_PORT))

print("UDP 服务端已启动，等待客户端连接...")

while True:
    # 接收数据报文和客户端地址
    data, client_addr = server_socket.recvfrom(BUFFER_SIZE)

    print("接收到来自客户端的数据：")
    print("客户端地址：", client_addr)
    print("数据内容：", data.decode())

server_socket.close()
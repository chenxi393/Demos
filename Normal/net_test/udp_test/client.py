import socket

UDP_IP = "127.0.0.1"
UDP_PORT = 8888
MESSAGE = "Hello, Server!"

# 创建 UDP 套接字
client_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

# 发送数据报文到服务器
client_socket.sendto(MESSAGE.encode(), (UDP_IP, UDP_PORT))

print("已发送数据到服务器：")
print("服务器地址：", UDP_IP)
print("服务器端口：", UDP_PORT)
print("发送内容：", MESSAGE)

client_socket.close()
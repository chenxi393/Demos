import http.client

# 创建一个 HTTP 连接
conn = http.client.HTTPConnection("localhost", 8000)

# 发送 GET 请求
conn.request("GET", "/")

# 获取服务器的响应
response = conn.getresponse()

# 读取响应内容
data = response.read()

# 打印响应内容
print(data.decode())

# 关闭连接
conn.close()
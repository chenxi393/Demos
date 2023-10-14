from http.server import HTTPServer, BaseHTTPRequestHandler

# 定义自定义的请求处理类
class MyHTTPRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # 设置响应状态码为 200 OK
        self.send_response(200)
        
        # 设置响应头的 Content-Type
        self.send_header('Content-type', 'text/plain')
        
        # 结束响应头的设置
        self.end_headers()
        
        # 构造响应内容
        response = "Hello, World!"
        
        # 将响应内容转换为字节流并发送给客户端
        self.wfile.write(response.encode())

# 定义服务器地址和端口号
server_address = ('', 8000)

# 创建 HTTP 服务器对象，并指定请求处理类
httpd = HTTPServer(server_address, MyHTTPRequestHandler)
print('服务器已启动，监听端口 8000...')

# 启动服务器，一直运行直到手动停止
httpd.serve_forever()
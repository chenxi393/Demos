#include <errno.h>
#include <netdb.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>

#define MAXLINE 1024

int main(int argc, char* argv[])
{

    // 1. 创建一个监听 socket
    int listenfd = socket(AF_INET, SOCK_STREAM, 0);
    if (listenfd < 0) {
        fprintf(stderr, "socket error : %s\n", strerror(errno));
        return -1;
    }

    // 2. 初始化服务器地址和端口
    struct sockaddr_in server_addr;
    bzero(&server_addr, sizeof(struct sockaddr_in));
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    server_addr.sin_port = htons(8888);

    // 3. 绑定地址+端口
    if (bind(listenfd, (struct sockaddr*)(&server_addr), sizeof(struct sockaddr)) < 0) {
        fprintf(stderr, "bind error:%s\n", strerror(errno));
        return -1;
    }

    printf("begin listen....\n");

    // 4. 开始监听
    if (listen(listenfd, 128)) {
        fprintf(stderr, "listen error:%s\n\a", strerror(errno));
        exit(1);
    }

    // 5. 获取已连接的socket
    struct sockaddr_in client_addr;
    socklen_t client_addrlen = sizeof(client_addr);
    int clientfd = accept(listenfd, (struct sockaddr*)&client_addr, &client_addrlen);
    if (clientfd < 0) {
        fprintf(stderr, "accept error:%s\n\a", strerror(errno));
        exit(1);
    }

    printf("accept success\n");

    char message[MAXLINE] = { 0 };

    while (1) {
        // 6. 读取客户端发送的数据
        int n = read(clientfd, message, MAXLINE);
        if (n < 0) { // 读取错误
            fprintf(stderr, "read error:%s\n\a", strerror(errno));
            break;
        } else if (n == 0) { // 返回 0 ，代表读到 FIN 报文
            fprintf(stderr, "client closed \n");

            int value = 1;
            // 下面这一句开启四次挥手 快速ACK
            if (setsockopt(clientfd, IPPROTO_TCP, TCP_QUICKACK, (char*)&value, sizeof(value)) < 0) {
                fprintf(stderr, "set TCP_QUCIK error:%s\n\a", strerror(errno));
            }
            close(clientfd); // 没有数据要发送，立马关闭连接
            break;
        }

        message[n] = 0;
        printf("received %d bytes: %s\n", n, message);
    }

    close(listenfd);
    return 0;
}
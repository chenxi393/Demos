#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <netdb.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <sys/socket.h>

int main(int argc, char *argv[])
{

    // 1. 创建一个监听 socket
    int connectfd = socket(AF_INET, SOCK_STREAM, 0);
    if(connectfd < 0)
    {
        fprintf(stderr, "socket error : %s\n", strerror(errno));
        return -1;
    }

    // 2. 初始化服务器地址和端口
    struct sockaddr_in server_addr;
    bzero(&server_addr, sizeof(struct sockaddr_in));
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    server_addr.sin_port = htons(8888);
    
    // 3. 连接服务器
    if(connect(connectfd, (struct sockaddr *)(&server_addr), sizeof(server_addr)) < 0)
    {
        fprintf(stderr,"connect error:%s\n", strerror(errno));
        return -1;
    }

    printf("connect success\n");


    char sendline[64] = "hello, i am xiaolin";

    //4. 发送数据
    int ret = send(connectfd, sendline, strlen(sendline), 0);
    if(ret != strlen(sendline)) {
        fprintf(stderr,"send data error:%s\n", strerror(errno));
        return -1;
    }

    printf("already send %d bytes\n", ret);

    sleep(1);

    //5. 关闭连接
    close(connectfd);
    return 0;
}
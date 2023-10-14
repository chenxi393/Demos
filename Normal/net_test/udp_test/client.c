#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

#define DEFAULT_PORT 8888
#define BUFFER_SIZE 1024

int main() {
    int clientSocket;
    struct sockaddr_in serverAddress;
    char buffer[BUFFER_SIZE];
    ssize_t bytesSent;

    // 创建 UDP 套接字
    clientSocket = socket(AF_INET, SOCK_DGRAM, 0);
    if (clientSocket < 0) {
        perror("Failed to create socket");
        exit(1);
    }

    // 设置服务器地址和端口号
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_addr.s_addr = inet_addr("127.0.0.1"); // 替换为实际的服务器 IP 地址
    serverAddress.sin_port = htons(DEFAULT_PORT);

    const char* message = "Hello, Server!";

    // 发送数据报文到服务器
    bytesSent = sendto(clientSocket, message, strlen(message), 0, (struct sockaddr*)&serverAddress, sizeof(serverAddress));
    if (bytesSent < 0) {
        perror("Failed to send data");
        exit(1);
    }

    printf("已发送数据到服务器：\n");
    printf("服务器地址：%s\n", inet_ntoa(serverAddress.sin_addr));
    printf("服务器端口：%d\n", ntohs(serverAddress.sin_port));
    printf("发送内容：%s\n", message);

    close(clientSocket);

    return 0;
}
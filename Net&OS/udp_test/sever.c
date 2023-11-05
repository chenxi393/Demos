#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

#define DEFAULT_PORT 8888
#define BUFFER_SIZE 1024

int main() {
    int serverSocket;
    struct sockaddr_in serverAddress, clientAddress;
    socklen_t clientAddressSize = sizeof(clientAddress);
    char buffer[BUFFER_SIZE];
    ssize_t bytesRead;

    // 创建 UDP 套接字
    serverSocket = socket(AF_INET, SOCK_DGRAM, 0);
    if (serverSocket < 0) {
        perror("Failed to create socket");
        exit(1);
    }

    // 绑定 IP 地址和端口号
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_addr.s_addr = INADDR_ANY;
    serverAddress.sin_port = htons(DEFAULT_PORT);
    if (bind(serverSocket, (struct sockaddr*)&serverAddress, sizeof(serverAddress)) < 0) {
        perror("Failed to bind socket");
        exit(1);
    }

    printf("UDP 服务端已启动，等待客户端连接...\n");

    while (1) {
        // 接收数据报文和客户端地址
        bytesRead = recvfrom(serverSocket, buffer, BUFFER_SIZE, 0, (struct sockaddr*)&clientAddress, &clientAddressSize);
        if (bytesRead < 0) {
            perror("Failed to receive data");
            exit(1);
        }

        printf("接收到来自客户端的数据：\n");
        printf("客户端地址：%s\n", inet_ntoa(clientAddress.sin_addr));
        printf("客户端端口：%d\n", ntohs(clientAddress.sin_port));
        printf("数据内容：%.*s\n", (int)bytesRead, buffer);
    }

    close(serverSocket);

    return 0;
}
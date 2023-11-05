#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>
#include <fcntl.h>
#include <unistd.h>

#define SHM_SIZE 1024

int main() {
    int fd = open("shared_memory_example.txt", O_RDWR); // 打开共享内存文件
    if (fd == -1) {
        perror("open");
        exit(1);
    }

    char *shared_memory = (char *)mmap(NULL, SHM_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (shared_memory == MAP_FAILED) {
        perror("mmap");
        exit(1);
    }

    printf("Message received: %s", shared_memory); // 从共享内存读取数据并打印

    munmap(shared_memory, SHM_SIZE); // 解除共享内存的映射
    close(fd); // 关闭共享内存文件

    return 0;
}
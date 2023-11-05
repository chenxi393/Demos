#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>
#include <fcntl.h>
#include <unistd.h>

#define SHM_SIZE 1024

int main() {
    int fd = open("shared_memory_example.txt", O_CREAT | O_RDWR, 0666); // 创建共享内存文件
    if (fd == -1) {
        perror("open");
        exit(1);
    }

    ftruncate(fd, SHM_SIZE); // 设置共享内存文件的大小

    char *shared_memory = (char *)mmap(NULL, SHM_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (shared_memory == MAP_FAILED) {
        perror("mmap");
        exit(1);
    }

    printf("Enter a message: ");
    fgets(shared_memory, SHM_SIZE, stdin); // 从标准输入读取数据到共享内存

    munmap(shared_memory, SHM_SIZE); // 解除共享内存的映射
    close(fd); // 关闭共享内存文件

    return 0;
}
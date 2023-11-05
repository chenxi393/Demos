#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>

#define SHM_SIZE 1024

int main() {
    key_t key = ftok("shared_memory_example", 1234); // 生成共享内存的键值
    int shmid = shmget(key, SHM_SIZE, 0666); // 获取共享内存段

    if (shmid == -1) {
        perror("shmget");
        exit(1);
    }

    char *shared_memory = (char *)shmat(shmid, NULL, 0); // 将共享内存连接到当前进程的地址空间
    if (shared_memory == (char *)-1) {
        perror("shmat");
        exit(1);
    }

    printf("Message received: %s", shared_memory); // 从共享内存读取数据并打印

    shmdt(shared_memory); // 分离共享内存

    shmctl(shmid, IPC_RMID, NULL); // 删除共享内存段

    return 0;
}
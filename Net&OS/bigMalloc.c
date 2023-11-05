#include <stdio.h>
#include <stdlib.h>

#define GB (1024 * 1024 * 1024)

int main() {
    size_t size = 10100000000;
    void *ptr = malloc(size);

    if (ptr == NULL) {
        printf("无法申请 %luGB 内存\n", size / GB);
    } else {
        printf("成功申请 %luGB 内存\n", size / GB);
        getchar();
        free(ptr);
    }

    return 0;
}
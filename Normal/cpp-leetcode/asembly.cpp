#include <stdio.h>

int main() {
    int a = 10;
    int b;

    asm("movl %0, %%eax;" //使用双百分号（%%）来表示单个百分号（%）
        "addl $5, %%eax;"
        "movl %%eax, %0;"
        : "=r" (b) // =r 表示是输出操作数 并分配一个通用寄存器
        : "r" (a) // 表示输入操作数
        : "%eax"); // 这里表示%eax会被适用

    printf("Result: %d\n", b);
    
    return 0;
}
#include <iostream>
using namespace std;
// 位元维度一面
// p
// /3 0 0 0 0 0| |0 0 -4|
class Allocator {
    int* arr;
    int size;
    int ptr;
    Allocator(int s)
    {
        size = s;
        ptr = 0;
        arr = new int(size);
        arr[ptr] = -1;
    }
    ~Allocator()
    {
        delete[] arr;
    }
    int allocate(int s)
    {
        // s=4
        if (ptr + s >= size) {
            // 分配失败
            return -1;
        }
        int t = ptr + s;
        arr[t] = ptr;
        // >=0表示未被使用
        // 负数表示使用过了
    }
};
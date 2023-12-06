#include <functional>
#include <iostream>
#include <pthread.h>
#include <thread>
using namespace std;

// 这个不知道 为什么死锁了 需要去学习多线程或者多进程的调试手段
// 看看为什么死锁了 就是打印不出东西 条件变量和互斥锁的实现
class ZeroEvenOdd {
private:
    int n;
    pthread_cond_t zero0;
    pthread_cond_t even0;
    pthread_cond_t odd0;
    pthread_mutex_t mm;
    int cur;
    bool flag;

public:
    ZeroEvenOdd(int n)
    {
        this->n = n;
        pthread_cond_init(&zero0, NULL);
        pthread_cond_init(&even0, NULL);
        pthread_cond_init(&odd0, NULL);
        pthread_mutex_init(&mm, NULL);
        cur = 1;
        flag = true;
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber)
    {
        while (true) {
            pthread_mutex_lock(&mm);
            while (flag == false) {
                pthread_cond_wait(&zero0, &mm);
            }
            if (cur > n) {
                pthread_mutex_unlock(&mm);
                return;
            }
            printNumber(0);
            printf("%d", 0);
            flag = false;
            if (cur % 2 == 1) {
                pthread_cond_signal(&odd0);
            } else {
                pthread_cond_signal(&even0);
            }
            pthread_mutex_unlock(&mm);
        }
    }

    void even(function<void(int)> printNumber)
    {
        while (true) {
            pthread_mutex_lock(&mm);
            while (cur % 2 == 1 || flag == true) {
                pthread_cond_wait(&even0, &mm);
            }
            printNumber(cur);
            printf("%d", cur);
            cur++;
            if (cur > n) {
                pthread_mutex_unlock(&mm);
                return;
            }
            flag = true;
            pthread_mutex_unlock(&mm);
            pthread_cond_signal(&zero0);
        }
    }

    void odd(function<void(int)> printNumber)
    {
        while (true) {
            pthread_mutex_lock(&mm);
            while (cur % 2 == 0 || flag == true) {
                pthread_cond_wait(&odd0, &mm);
            }
            printNumber(cur);
            printf("%d", cur);
            cur++;
            if (cur > n) {
                return;
            }
            flag = true;
            pthread_mutex_unlock(&mm);
            pthread_cond_signal(&zero0);
        }
    }
};
void print(int i)
{
    std::cout << i;
}

int main()
{
    ZeroEvenOdd zeroEvenOdd(10);

    std::thread t1(&ZeroEvenOdd::zero, &zeroEvenOdd, print);
    std::thread t2(&ZeroEvenOdd::even, &zeroEvenOdd, print);
    std::thread t3(&ZeroEvenOdd::odd, &zeroEvenOdd, print);

    t1.join();
    t2.join();
    t3.join();
}

#include <functional>
#include <iostream>
#include <semaphore.h>
#include <thread>
using namespace std;
class ZeroEvenOdd {
private:
    int n;
    sem_t printOdd, printEven, numDone;

public:
    ZeroEvenOdd(int n)
    {
        this->n = n;
        sem_init(&printOdd, 0, 0);
        sem_init(&printEven, 0, 0);
        sem_init(&numDone, 0, 1);
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber)
    {
        for (int i = 1; i <= n; ++i) {
            sem_wait(&numDone);
            printNumber(0);
            if (i % 2 == 0) {
                sem_post(&printEven);
            } else {
                sem_post(&printOdd);
            }
        }
    }

    void even(function<void(int)> printNumber)
    {
        for (int i = 2; i <= n; i += 2) {
            sem_wait(&printEven);
            printNumber(i);
            sem_post(&numDone);
        }
    }

    void odd(function<void(int)> printNumber)
    {
        for (int i = 1; i <= n; i += 2) {
            sem_wait(&printOdd);
            printNumber(i);
            sem_post(&numDone);
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

#include <chrono>
#include <iostream>

using namespace std;
using namespace std::chrono;

double fib(double n)
{
    if (n <= 1)
        return n;
    return fib(n - 1) + fib(n - 2);
}

int main()
{
    double n = 42; // 要计算的斐波那契数列的项数

    // 记录开始时间
    auto start = high_resolution_clock::now();

    // 计算斐波那契数列
    double result = fib(n);

    // 记录结束时间
    auto end = high_resolution_clock::now();

    // 计算时间差（单位：秒）
    duration<double> time_span = duration_cast<duration<double>>(end - start);
    double time_taken = time_span.count();


    cout << "计算耗时: " << time_taken << " 秒" << endl;

    return 0;
}
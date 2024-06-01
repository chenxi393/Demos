#include <iostream>
#include <queue>
using namespace std;
class MedianFinder {
public:
    // 两个堆 实现简单多了
    // addNum 时间复杂度 O(logn) findMedian() O(1)
    // 使用两个堆保存数组大的一部分和小的一部分
    priority_queue<int, vector<int>, greater<int>> A; // 小顶堆，保存较大的一半
    priority_queue<int, vector<int>, less<int>> B; // 大顶堆，保存较小的一半
    /** initialize your data structure here. */
    MedianFinder()
    {
    }

    void addNum(int num)
    {
        if (A.size() != B.size()) { // A的元素多 弹出A最小的 B++
            A.push(num);
            B.push(A.top());
            A.pop();
        } else { // 偶数个 元素放入B 然后弹除最大 给A A++
            B.push(num);
            A.push(B.top());
            B.pop();
        }
    }

    double findMedian()
    {
        if (A.size() == B.size()) { // 偶数个 A的最小
            return (A.top() + B.top()) / 2.0;
        } else {
            return A.top();
        }
    }
};

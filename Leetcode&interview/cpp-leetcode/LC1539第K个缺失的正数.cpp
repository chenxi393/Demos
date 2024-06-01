#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    // arr = [2,3,4,7,11], k = 5
    int findKthPositive(vector<int>& arr, int k)
    {
        int left = 0;
        int right = arr.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            // arr[mid]-mid -1 也就当前下标 缺失的元素个数
            if (arr[mid] - mid - 1 < k) {
                left = mid + 1;
            } else {
                // 出口只在这 那么left一定是正确答案
                // 这算是二分的一种形式
                right = mid - 1;
            }
        }
        return left + k;
    }
};

int main()
{

    return 0;
}

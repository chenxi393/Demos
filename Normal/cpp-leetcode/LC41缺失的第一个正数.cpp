#include <iostream>

using namespace std;

class Solution {
public:
    // 原理交换法 也可以三次遍历 标记法 之类的
    int firstMissingPositive(vector<int>& nums)
    {
        for (int i = 0; i < nums.size(); i++) {
            // 当前的数不等于 这个位置的数
            // 当前数需要在 1-N之间
            while (nums[i] <= nums.size() && nums[i] >= 1 && nums[i] - 1 != i) {
                // 对应位置已经安排好 则跳出
                if (nums[i] == nums[nums[i] - 1]) {
                    break;
                }
                swap(nums[i], nums[nums[i] - 1]);
            }
        }
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] != i + 1) {
                return i + 1;
            }
        }
        return nums.size() + 1;
    }
};

int main()
{

    return 0;
}

#include <iostream>
#include <vector>
using namespace std;

// [2,1,4,3]  2  3
// 1
// [2,9,2,5,6]  2  8
// 一个合理区间n的子区间个数就是n! 所有遍历合理区间可以得到
// 1 + 2 + 3 = 3!
class Solution {
public:
    // 此题可以转换为求 所有小于等于right的子区间 - 所有小于left的子区间
    int numSubarrayBoundedMax(vector<int>& nums, int left, int right)
    {
        return Count(nums, right) - Count(nums, left - 1);
    }

    int Count(vector<int>& nums, int B)
    {
        // 遍历区间累计可以转化为阶乘数（也就是子区间数）
        int cnt = 0;
        int ans = 0;
        for (auto i : nums) {
            if (i <= B) {
                cnt++;
            } else {
                cnt = 0;
            }
            ans += cnt;
        }
        return ans;
    }
};

class Solution1 {
public:
    // 法二 一次遍历
    int numSubarrayBoundedMax(vector<int>& nums, int left, int right)
    {
        int last1 = -1;
        int last2 = -1;
        int ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] >= left && nums[i] <= right) {
                // 当前元素在区间内
                // 当前下标 要是当前下标为0 0-（-1） 刚好是1
                // 这也是为什么初始化为-1
                last1 = i;
                ans += (last1 - last2);
            } else if (nums[i] > right) {
                // 超过区间了
                last2 = i;
                last1 = -1;
                // 重置区间
            }else if(last1!=-1){
                ans += (last1 - last2);
            }
             
        }
        return ans;
    }
};

int main()
{

    return 0;
}

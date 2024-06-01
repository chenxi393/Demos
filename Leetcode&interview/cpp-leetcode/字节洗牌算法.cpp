#include <iostream>
#include <random>
using namespace std;

// 对K个用户进行 抽奖 中间概论公平 洗牌算法？？？ 
// 空间复杂度O(1) 推导一下如何保证概率公平     

// 如果抽一个人就 random 就行
// 注意场景 是抽k个人 用random抽第一个人 再抽肯定不是等概率的
vector<int> GetRandomN(vector<int>& nums, int k)
{
    for (int i = nums.size() - 1; i >= 0; i--) {
        int temp = rand() % (i + 1);
        cout << "temp:" << temp << endl;
        swap(nums[temp], nums[i]);
    }
    return vector<int>(nums.begin(), nums.begin() + k);
}

int main()
{
    vector<int> nums;
    for (int i = 0; i < 10; i++) {
        nums.push_back(i + 1);
    }
    auto ans = GetRandomN(nums, 5);
    for (auto i : ans) {
        cout << i << endl;
    }
    return 0;
}

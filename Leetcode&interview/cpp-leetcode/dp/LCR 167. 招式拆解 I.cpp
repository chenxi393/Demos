#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
public:
    // 看了解析的 感觉自己还是想不到啊
    // 滑动窗口 + map
    int dismantlingAction(string arr)
    {
        unordered_map<char, int> mp;
        int i = -1, res = 0;
        for (int j = 0; j < arr.length(); j++) {
            // mp 里维护字符最新的下标
            if (mp.find(arr[j]) != mp.end()) {
                // 说明找到重复的数了
                // 如果重复的数 索引大于i 则更新 否则不更新 无影响
                i = max(i, mp[arr[j]]);
            }
            // 更新mp最新
            mp[arr[j]] = j;
            res = max(res, j - i);
        }
        return res;
    }
};

int main()
{

    return 0;
}

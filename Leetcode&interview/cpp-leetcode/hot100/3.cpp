#include <algorithm>
#include <iostream>
#include <unordered_map>
using namespace std;

// 这题 改造一下 变成返回最长子串 而不是长度
class Solution {
public:
    int lengthOfLongestSubstring(string s)
    {
        unordered_map<int, int> mp;
        int i = -1, j = 0;
        int res = 0;
        while (j < s.length()) {
            auto t = mp.find(s[j]);
            if (t == mp.end() || t->second <= i) {
                res = max(res, j - i);
            } else {
                i = t->second;
            }
            mp[s[j]] = j;
            j++;
        }
        return res;
    }

    string lengthOfLongestSubstring2(string s)
    {
        unordered_map<int, int> mp;
        int i = -1, j = 0;
        int res = 0;
        int resl = 0;
        while (j < s.length()) {
            auto t = mp.find(s[j]);
            if (t == mp.end() || t->second <= i) {
                res = max(res, j - i);
                resl = i + 1;
            } else {
                i = t->second;
            }
            mp[s[j]] = j;
            j++;
        }
        return s.substr(resl, res);
    }
};

int main()
{

    return 0;
}

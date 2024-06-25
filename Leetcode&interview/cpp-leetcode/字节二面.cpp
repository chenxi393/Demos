#include <algorithm>
#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

// 自己面试胡乱写的（暴力） 感觉面试官不懂C++ 才。。过的
// 这是自己面试写的 只是凑巧样例过了 出大问题...
// abcabcab 这是面试的样例
int LongestSubstr(string s)
{
    // 先找出所有的字串
    unordered_set<string> substr;
    for (int i = 0; i < s.length(); i++) {
        for (int j = 0; j < s.length(); j++) {
            // 这里的api写错了 这里应该是j-i+1 66 面试的时候写错了
            // FIXME 一个是这个写错了
            // string sss = s.substr(i, j); //面试的写法
            string sss = s.substr(i, j - i + 1);
            if (substr.find(sss) == substr.end()) {
                substr.insert(sss);
            }
        }
    }
    vector<string> ans;
    for (auto sss : substr) {
        unordered_set<char> unique;
        int i = 0;
        for (; i < sss.length(); i++) {
            if (unique.find(sss[i]) != unique.end()) {
                break;
            } else {
                unique.insert(sss[i]);
            }
        }
        if (i == sss.length()) {
            ans.push_back(sss);
        }
    }
    // FIXME 一个这里不应该sort sort完全是按照字典序
    sort(ans.begin(), ans.end());
    for (auto sss : ans) {
        cout << sss << endl;
    }
    return ans[ans.size() - 1].length();
}

// 自己复盘重写的暴力 修复了面试的bug
// 剪枝优化后 过不了最后一个测试样例
int LongestSubstrFix(string s)
{
    // 先找出所有的字串
    unordered_set<string> substr;
    int maxSubStr = 0;
    for (int i = 0; i < s.length(); i++) {
        for (int j = 0; j < s.length(); j++) {
            if (j - i + 1 < maxSubStr) {
                continue;
            }
            string sss = s.substr(i, j - i + 1);
            if (substr.find(sss) == substr.end()) {
                substr.insert(sss);
                int i = 0;
                unordered_set<char> unique;
                for (; i < sss.length(); i++) {
                    if (unique.find(sss[i]) != unique.end()) {
                        break;
                    } else {
                        unique.insert(sss[i]);
                    }
                }
                if (i == sss.length()) {
                    maxSubStr = max(maxSubStr, int(sss.length()));
                }
            }
        }
    }
    return maxSubStr;
}

// 正确O(N) 的解法 使用滑动窗口
// 虽然过OJ了 但是时间很长
int lengthOfLongestSubstring(string s)
{
    int maxSubLength = 0;
    for (int i = 0; i < s.length(); i++) {
        int j = i + 1;
        unordered_set<char> unique;
        unique.insert(s[i]);
        while (j < s.length()) {
            if (unique.find(s[j]) == unique.end()) {
                unique.insert(s[j]);
            } else {
                break;
            }
            j++;
        }
        maxSubLength = max(maxSubLength, j - i);
    }
    return maxSubLength;
}

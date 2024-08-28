#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
    vector<string> letterCombinations(string digits)
    {
        if (digits.length() == 0) {
            return {};
        }
        unordered_map<char, string> mp;
        mp['2'] = "abc";
        mp['3'] = "def";
        mp['4'] = "ghi";
        mp['5'] = "jkl";
        mp['6'] = "mno";
        mp['7'] = "pqrs";
        mp['8'] = "tuv";
        mp['9'] = "wxyz";
        vector<string> ans;
        dfs(ans, mp, 0, digits, "");
        return ans;
    }

    void dfs(vector<string>& ans, unordered_map<char, string>& mp, int i, string digits, string now)
    {
        if (i == digits.length()) {
            ans.push_back(now);
            return;
        }

        for (char v : mp[digits[i]]) {
            now += v;
            // 其实可以直接函数传参 因为是值传递
            dfs(ans, mp, i + 1, digits, now);
            now = now.substr(0, now.length() - 1);
        }
    }
};

int main()
{

    return 0;
}

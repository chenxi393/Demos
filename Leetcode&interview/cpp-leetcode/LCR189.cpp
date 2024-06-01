#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    vector<string> goodsOrder(string goods)
    {
        vector<string> ans;
        dfs(0, goods, ans);
        return ans;
    }

    void dfs(int k, string& s, vector<string>& ans)
    {
        if (k == s.length()) {
            ans.push_back(s);
            return;
        }
        // 碰见全排列好几次了 但是总是有点写不对
        // 还要去重 字符串里会有相同元素
        for (int i = k ; i < s.length(); i++) {
            swap(s[i], s[k]);
            dfs(k+1, s, ans);// 注意这里 是固定前一位
            swap(s[i], s[k]);
        }
    }
};

int main()
{
    Solution s;
    vector<string> ans = s.goodsOrder("agew");
    for (auto i : ans) {
        cout << i << endl;
    }
}
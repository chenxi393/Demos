#include <iostream>
#include <set>
#include <vector>
// 其实就是全排列去重（）
// B站一面还简单一点 但是自己写的有问题
// 还是得多做题。。。。
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
        set<char> ss;
        // 因为固定了这一位，后面的顺序都无所谓
        // 反正最终都会全排列
        for (int i = k; i < s.length(); i++) {
            if (ss.find(s[i]) != ss.end()) {
                continue;
            }
            ss.insert(s[i]);
            swap(s[i], s[k]);
            dfs(k + 1, s, ans); // 注意这里 是固定前一位
            swap(s[i], s[k]);
        }
    }
};

int main(int argc, char* argv[])
{
    Solution s;
    vector<string> ans;

    ans = s.goodsOrder("agew");
    for (auto item : ans) {
        cout << item << " ";
    }
    cout << endl;

    return 0;
}

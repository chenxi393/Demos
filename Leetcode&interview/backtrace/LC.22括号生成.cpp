#include <bits/stdc++.h>
#include <iostream>
#include <stack>
#include <vector>

using namespace std;

// ATT还有优化空间 可能压根不需要
// 写的有点复杂了
class Solution {
public:
    vector<string> ans;
    vector<string> generateParenthesis(int n)
    {
        // 两种选择 左括号 还是右括号
        // 左右括号数量只有 n个
        // 要判断是否有效
        sum[0] = sum[1] = n;
        dfs(n, 0, "");
        return ans;
    }

    char sss[2] = { '(', ')' };
    int sum[2] = {};
    void dfs(int n, int k, string cur)
    {
        if (k >= n * 2) {
            if (valid(cur))
                ans.push_back(cur);
            // 可能需要判断括号是否有效
            return;
        }

        for (int i = 0; i < 2; i++) {
            // T可以用括号数量 来进行剪枝？
            if (sum[i] <= 0) {
                continue;
            }
            sum[i]--;
            cur.push_back(sss[i]);
            dfs(n, k + 1, cur);
            cur.pop_back();
            sum[i]++;
        }
    }

    bool valid(string cur)
    {
        // "((())"
        stack<char> s;
        for (auto v : cur) {
            // 栈里 是不可能右括号的
            if (v == '(') {
                s.push(v);
            } else {
                if (s.empty()) {
                    return false;
                }
                s.pop();
            }
        }
        return s.empty();
    }
};

int main()
{
    int n;
    cin >> n;
    Solution* s = new Solution();
    auto ans = s->generateParenthesis(n);

    for (auto v : ans) {
        cout << v << endl;
    }

    return 0;
}

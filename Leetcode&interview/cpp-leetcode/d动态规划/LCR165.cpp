#include <iostream>
#include <vector>
using namespace std;

// 这个题 不能因为放在动态规划里面就想动态规划
// 其实应该先想搜索
class Solution {
public:
    int count = 0;
    int crackNumber(int ciphertext)
    {
        // 先拆成字符串
        vector<char> ans;
        // 注意0 这种特殊情况
        if (ciphertext == 0) {
            return 1;
        }
        while (ciphertext > 0) {
            ans.push_back(ciphertext % 10);
            ciphertext /= 10;
        }
        ans.push_back(0);
        vector<char> temp(ans.rbegin(), ans.rend());
        dfs(false, 1, temp);
        return count;
    }

    void dfs(bool former, int pos, const vector<char>& ans)
    {
        // 不满足就退出
        if (former && (ans[pos - 1] == 0 || ans[pos - 1] * 10 + ans[pos] > 25)) {
            return;
        }
        if (pos == ans.size() - 1) {
            count++;
            return;
        }
        // 注意 如果这个是 true 那下一个就不能是true 不能被用两次
        if (former) {
            dfs(false, pos + 1, ans);
        } else {
            dfs(false, pos + 1, ans);
            dfs(true, pos + 1, ans);
        }
    }
};

class Solution1 {
public:
    // 纯递归 参考题解的 然后自己写 感觉很简洁
    int crackNumber(int ciphertext)
    {
        if (ciphertext <= 9) {
            return 1;
        }
        int num = ciphertext % 100;
        if (num <= 9 || num >= 26) {
            // 翻译不了两位
            return crackNumber(num / 10);
        }
        // 可以两位翻译 则有两种方式
        return crackNumber(num / 10) + crackNumber(num / 100);
    }
};

class Solution2 {
public:
    // 纯递归 参考题解的 然后自己写 感觉很简洁
    int crackNumber(int ciphertext)
    {
        // 递归树
        string num = "";
        num.(ciphertext);

    }
};

int main()
{
    int sss = 216612;
    Solution s;
    cout << s.crackNumber(sss);

    return 0;
}

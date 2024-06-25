#include <algorithm>
#include <iostream>
using namespace std;

class Solution {
public:
    // 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
    // 丑数是只包含质因数 2、3、5 的正整数；1 是丑数
    int nthUglyNumber(int n)
    {
        int dp[n];
        int a = 0, b = 0, c = 0;
        for (int i = 1; i < n; i++) {
            // 丑数 只能由丑数乘以2 3 5得到
            dp[i] = min(dp[a] * 2, dp[b] * 3);
            dp[i] = min(dp[i], dp[c] * 5);
            // 计算当前递增的最小丑数
            // 当当前丑数可以被2 3 5 计算出来时，则对应索引++
            if (dp[i] == dp[a] * 2) {
                a++;
            }
            if (dp[i] == dp[b] * 3) {
                b++;
            }
            if (dp[i] == dp[c] * 5) {
                c++;
            }
        }
        return dp[n - 1];
    }
};

int main()
{

    return 0;
}

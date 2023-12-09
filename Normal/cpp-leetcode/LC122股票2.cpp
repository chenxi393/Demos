#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    // 可能还需要再做做
    // 解法一：其实就是两相邻数 若递增 则累加 差值
    // 解法二：动态规划
    // 最初的想法：DFS 搜索 会超时 考虑 持有->卖出 持有->持有
    // 非持有->买入 非持有->非持有 四种状态
    // dp[i][j] i表示天数 j=0表示当天持有的现金 j=1表示持有股票

    // 贪心法
    int maxProfit(vector<int>& prices)
    {
        int buy = INT32_MAX;
        int profit = 0;
        for (auto p : prices) {
            if (p - buy > 0) {
                profit += (p - buy);
            }
            buy = p;
        }
        return profit;
    }
    // 动态规划法
    int maxProfitDP(vector<int>& prices)
    {
        int dp[prices.size()][2];
        dp[0][0] = 0; // 本金
        dp[0][1] = -prices[0]; // 假设买入了第一条的股票
        for (int i = 1; i < prices.size(); i++) {
            // 如果当前价格卖出能赚的
            int cur = prices[i] + dp[i - 1][1];
            // 卖出 比历史最优更大 则卖出 更新历史最优
            if (cur > dp[i - 1][0]) {
                dp[i][0] = cur;
            } else {
                dp[i][0] = dp[i - 1][0];
            }
            // 等价于 dp[i][0] =max(dp[i-1][0], dp[i-1][1]+prices[i])
            cur = dp[i - 1][0] - prices[i];
            // 如果当前价格比历史买入更低 则更新买入价格
            if (cur > dp[i - 1][1]) {
                dp[i][1] = cur;
            } else {
                dp[i][1] = dp[i - 1][1];
            }
            // 等价于 dp[i][1] = max(dp[i-1][1],dp[i-i][0]-prices[i])
        }
        return dp[prices.size()-1][0];
    }
};

int main()
{

    return 0;
}

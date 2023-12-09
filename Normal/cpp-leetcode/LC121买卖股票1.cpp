#include <iostream>
#include <vector>
using namespace std;

/*
dp[i] 表示前 i 天的最大利润，因为我们始终要使利润最大化，则：
dp[i]=max(dp[i−1],prices[i]−minprice)
*/
class Solution {
public:
    int maxProfit(vector<int>& prices)
    {
        int minbuy = INT32_MAX;
        int maxsell = 0;
        int temp;
        for (int i = 0; i < prices.size(); i++) {
            temp = prices[i] - minbuy;
            maxsell = max(temp, maxsell);
            if (minbuy > prices[i]) {
                // 找到历史最低点买入
                minbuy = prices[i];
            }
        }
        return maxsell;
    }
};

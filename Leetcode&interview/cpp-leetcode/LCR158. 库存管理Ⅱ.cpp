#include <iostream>
#include <vector>
using namespace std;

// 简言之 这题就是O(n) 找到数组里 一半以上的数
class Solution {
public:
    // 1 <= stock.length <= 50000
    int inventoryManagement(vector<int>& stock)
    {
        // 4 5 2 3 2 3 2
        int cur = stock[0];
        int score = 0;
        for (int i = 0; i < stock.size(); i++) {
            if (score == 0) {
                cur = stock[i];
            }
            if (cur == stock[i]) {
                score++;
            } else {
                score--;
            }
        }
        return cur;
    }
    // 注意如果题目 没有满足条件的数
    // 需要遍历原数组 判断
};
int main()
{

    return 0;
}

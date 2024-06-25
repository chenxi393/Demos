#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

// 自己写的 17min
class Solution {
public:
    int maxSales(vector<int>& sales)
    {
        int* max_ = new int[sales.size() + 1] { 0 };
        int cur = 0;
        int max__ = 0;
        for (int i = 0; i < sales.size(); i++) {
            if (cur + sales[i] >= 0) {
                cur += sales[i];
            }
            max__ = max(max__, cur);
        }
        return max__;
    }
};

// 参考 感觉这样逻辑更清晰一点
class Solution1 {
public:
    int maxSales(vector<int>& sales)
    {
        int max__ = sales[0]; // 初始化最大为 第一个元素
        for (int i = 1; i < sales.size(); i++) {
            if (sales[i - 1] > 0) {
                // 前一个元素对 最大有贡献
                sales[i] += sales[i - 1];
            }
            // 否则就是没有贡献
            // 两种情况都取最大
            max__ = max(sales[i], max__);
        }
        return max__;
    }
};
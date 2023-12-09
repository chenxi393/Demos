#include <iostream>
#include <vector>
using namespace std;
// 全排列 11/29 B站SRE 一面 有点没做出来
// 输入 1 2 3
// 【1 2 3】【1 3 2】【2 1 3】
// 【2 3 1】【3 2 1】【3 1 2】 
// 全排列 123 的结果数就是 n！ n的阶乘
// 每次固定一位 固定第一位 有3种选择 循环3次 固定第二位两种...
void func(vector<int>& num, int k);
vector<vector<int>> ans;
int main()
{
    vector<int> nums;
    nums.push_back(1);
    nums.push_back(2);
    nums.push_back(3);
    func(nums, 0);
    for (int i = 0; i < ans.size(); i++) {
        for (int j = 0; j < ans[i].size(); j++) {
            cout << ans[i][j] << " ";
        }
        cout << endl;
    }
}

void func(vector<int>& num, int k)
{
    if (k == num.size()) {
        ans.push_back(num);
        return;
    }
    for (int i = k; i < num.size(); i++) {
        swap(num[i], num[k]);
        // 应该传递的是k+1 而不是i+1 6 这都没写对 
        func(num, i + 1);
        swap(num[i], num[k]);
    }
}

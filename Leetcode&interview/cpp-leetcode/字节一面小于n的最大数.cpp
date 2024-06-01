#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

// 12/25 一面问的问题 这个是面试的时候写的
int maxN(const vector<int>& nums, int n)
{
    // 小于n
    n -= 1;
    // 将 nums 放入0-9的
    bool isExsit[10] = { false };
    int maxNum = 0;
    for (int i = 0; i < nums.size(); i++) {
        isExsit[nums[i]] = true;
        maxNum = max(maxNum, nums[i]);
    }
    vector<int> ans;
    while (n > 0) {
        ans.push_back(n % 10);
        n /= 10;
    }
    reverse(ans.begin(), ans.end());
    // 6 这里面试还写错了 .....
    // reverse(nums.begin(), nums.end()); 面试写的
    // 将 n 拆开
    // 从高位遍历到低位
    for (int i = 0; i < ans.size(); i++) {
        // 一旦有不在里的位
        if (!isExsit[ans[i]]) {
            while (ans[i] > 0 && !isExsit[ans[i]]) {
                ans[i]--;
            }
            // 将后面的位都置成最大 6 这里也写错。。。 面试的时候 k 写成i
            for (int k = i + 1; k < ans.size(); k++) {
                ans[k] = maxNum;
            }
            break;
        }
    }
    int res = 0;
    for (int i = 0; i < ans.size(); i++) {
        res = res * 10 + ans[i];
    }
    return res;
}

// k表示当前遍历n的位 n是被拆分的数字 nums是升序
// 这个复盘重新写的
int maxNByDFS(const vector<int>& nums, vector<int>& ans, int k)
{
    if (ans.size() == k) {
        int res = 0;
        for (auto i : ans) {
            res = res * 10 + i;
        }
        return res;
    }
    // 一旦有不在里的位
    int j = nums.size() - 1;
    while (j >= 0 && nums[j] > ans[k]) {
        j--;
    }
    // 说明num所有位都大于该位
    if (j < 0) { // 说明所有位都大于这个数
        if (k == 0) { // 如果k==0 将此位置0 进行下一位
            // 当前位为空
            // 后面的位全取大
            ans[k] = 0;
            for (int i = k + 1; i < ans.size(); i++) {
                ans[i] = nums[nums.size() - 1];
            }
            k = ans.size();
        } else {
            // 回退到前一位
            // 找到小于的
            // 手动的把前一位减小？？
            ans[k - 1] -= 1;
            // FIXME 需要把这一位以及后面的位全都置最大
            // 因为 特殊样例1
            for (int i = k; i < ans.size(); i++) {
                ans[i] = nums[nums.size() - 1];
            }
            k = k - 1;
        }
    } else if (nums[j] < ans[k]) {
        // 当前位小于 它 则填充这一位 后面全取大
        ans[k] = nums[j];
        for (int i = k + 1; i < ans.size(); i++) {
            ans[i] = nums[nums.size() - 1];
        }
        k = ans.size();
    } else if (nums[j] == ans[k]) {
        k++;
    }
    return maxNByDFS(nums, ans, k);
}

int main()
{
    // 23121  // 2 4 9
    // 23132  // 2 4 9
    // 21132  // 2 4 9 // 面试写的这个过不了
    // 24132  // 2 4 9
    // 24132  // 1 4 9
    // 14132  // 2 4 9
    // 14132  // 1 4 9
    // 54169  // 2 4 5 8 // 面试写的这个过不了
    // 54922  // 6 7 8 9
    // 66922  // 6 7 8 9  // 特殊样例1
    int n = 66922;
    vector<int> nums;
    nums.push_back(6);
    nums.push_back(7);
    nums.push_back(8);
    nums.push_back(9);
    cout << maxN(nums, n) << endl;
    vector<int> N;
    while (n > 0) {
        N.push_back(n % 10);
        n /= 10;
    }
    reverse(N.begin(), N.end());
    cout << maxNByDFS(nums, N, 0) << endl;

    return 0;
}
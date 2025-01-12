#include <bits/stdc++.h>
#include <iostream>
#include <stack>
#include <vector>

using namespace std;

// https://www.lanqiao.cn/problems/89/learning/
// 蓝桥杯 2016国赛

// DFS 左上角到右下角 由箭数 给出唯一路径数

// ATT
// 1. 写的有点繁琐了 看下怎么优化
// 2. 剪枝
// 3. 时间复杂度 4^(n^2) 每步有4个方向选择，一共可能有n^2步
// 4. （额外优化方法）比如如果当前列要变为0，检测前面的列是否都是0，
// 可以这样优化：只记录上一次变成0的列是谁，
// 如果是即将变成0的列编号-1，则合法，否则是非法

// 1. 确定要存储的状态
int n;
vector<vector<bool>> is_visited;
vector<int> north;
vector<int> west;
vector<int> path; //  输出路径 0 1 2 3

int dx[4] = { 1, 0, -1, 0 };
int dy[4] = { 0, 1, 0, -1 };

bool dfs(int x, int y)
{
    // 2. 判断越界
    if (x < 0 || y < 0 || x >= n || y >= n) {
        return false;
    }

    if (north[x] == 0 || west[y] == 0) {
        return false;
    }

    if (is_visited[x][y]) {
        return false;
    }

    // 7. 剪枝 即将要走的这个点 == 1 但是 前面的点 不为0 说明要回溯 但是回溯不了
    if (north[x] == 1 && accumulate(north.begin(), north.begin() + x, 0) != 0)
        return false;
    if (west[y] == 1 && accumulate(west.begin(), west.begin() + y, 0) != 0)
        return false;

    // 3. 更新状态
    is_visited[x][y] = true;
    north[x]--;
    west[y]--;
    path.push_back(x + y * n);

    // 4. 结束状态 成功达到右下角 仅一种情况
    if (x == n - 1 && y == n - 1) {
        // 需要检测是否箭数都为0
        if (accumulate(north.begin(), north.end(), 0) == 0 && accumulate(west.begin(), west.end(), 0) == 0) {
            return true;
        }
    }

    // 5. 状态转移 遍历四个点
    for (int i = 0; i < 4; i++) {
        // ATT 这里只有一种结果 如果拿到结果 就返回 不再搜索
        if (dfs(x + dx[i], y + dy[i])) {
            return true;
        }
    }

    // 6. 还原状态
    is_visited[x][y] = false;
    north[x]++;
    west[y]++;
    path.pop_back();
    return false;
}

void print(vector<int> nums)
{
    for (int i = 0; i < nums.size(); i++) {
        cout << nums[i];
        if (i != nums.size() - 1) {
            cout << " ";
        }
    }
    cout << endl;
}

void init()
{
    int t;
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> t;
        north.push_back(t);
    }
    for (int i = 0; i < n; i++) {
        cin >> t;
        west.push_back(t);
    }
    for (int i = 0; i < n; i++) {
        vector<bool> t;
        for (int j = 0; j < n; j++) {
            t.push_back(false);
        }
        is_visited.push_back(t);
    }
}

int main()
{
    init();
    dfs(0, 0);
    print(path);
    return 0;
}
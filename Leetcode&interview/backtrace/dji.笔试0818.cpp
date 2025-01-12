#include <iostream>
#include <vector>

using namespace std;

/*
无人机前进的方向要与上次保持一致 初始y+1 向下 不可达 则顺时针旋转90°


0 1 0
0 0 1
0 0 0

3(不是6也不是5)
*/

class Solution {
public:
    /* Write Code Here */
    int ans = 1, n = 0, m = 0;
    vector<vector<int>> block; // 输入的屏障
    vector<vector<vector<bool>>> direction_visited; // 保存该点，某方向 是否访问过了
    vector<vector<bool>> pos_visited; // 保存该点是否访问过了

    vector<int> dx = { 0, 1, 0, -1 }, dy = { 1, 0, -1, 0 };

    void dfs(int x, int y, int d)
    {
        for (int i = 0; i < 4; i++) {
            int nd = (d + i) % 4; // 方向要和上一个方向保持一致
            int nx = x + dx[nd], ny = y + dy[nd];
            if (nx > -1 && nx < n && ny > -1 && ny < m && !block[nx][ny]) { // 合法坐标，并且可访问
                if (!pos_visited[nx][ny])
                    ans++; // 统计过了 就不需要统计了
                pos_visited[nx][ny] = true;
                if (!direction_visited[x][y][nd]) { // 当前方向没有被访问过，防止死循环
                    direction_visited[x][y][nd] = true;
                    dfs(nx, ny, nd);
                }
                return;
            }
        }
    }

    int numberOfPatrolBlocks(vector<vector<int>> input)
    {
        n = input.size(), m = input[0].size();
        block = input;
        pos_visited = vector<vector<bool>>(n, vector<bool>(m));
        direction_visited = vector<vector<vector<bool>>>(n, vector<vector<bool>>(m, vector<bool>(4)));
        ans = 1;
        pos_visited[0][0] = true;
        dfs(0, 0, 0);
        return ans;
    }
};

int main()
{
    int res;

    int block_rows = 0;
    int block_cols = 0;
    cin >> block_rows >> block_cols;
    vector<vector<int>> block(block_rows);
    for (int block_i = 0; block_i < block_rows; block_i++) {
        for (int block_j = 0; block_j < block_cols; block_j++) {
            int block_tmp;
            cin >> block_tmp;
            block[block_i].push_back(block_tmp);
        }
    }
    Solution* s = new Solution();
    res = s->numberOfPatrolBlocks(block);

    cout << res << endl;

    return 0;
}

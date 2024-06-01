#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    bool exist(vector<vector<char>>& board, string word)
    {
        vector<vector<bool>> visited;
        visited.resize(board.size());
        for (int i = 0; i < visited.size(); i++) {
            visited[i].resize(board[i].size());
        }
        for (int i = 0; i < board.size(); i++) {
            for (int j = 0; j < board[i].size(); j++) {
                if (word[0] == board[i][j] && dfs(board, visited, i, j, word, 0)) {
                    return true;
                }
            }
        }
        return false;
    }

    bool dfs(vector<vector<char>>& board, vector<vector<bool>>& visited, int i, int j, const string& word, int k)
    {
        visited[i][j] = true;
        k++;
        if (k == word.size()) {
            return true;
        }
        int p[4][2] = { { 1, 0 }, { -1, 0 }, { 0, 1 }, { 0, -1 } };
        for (int p1 = 0; p1 < 4; p1++) {
            int curi = i + p[p1][0];
            int curj = j + p[p1][1];
            if (curi < board.size() && curj < board[i].size() && curi >= 0 && curj >= 0) {
                if (!visited[curi][curj] && word[k] == board[curi][curj]) {
                    if (dfs(board, visited, curi, curj, word, k)) {
                        return true;
                    }
                }
            }
        }
        visited[i][j] = false;
        return false;
    }
};

int main()
{
    vector<vector<char>> board = {
        { 'A', 'B', 'C', 'E' },
        { 'S', 'F', 'C', 'S' },
        { 'A', 'D', 'E', 'E' }
    };
    string word = "ABCCED";
    Solution s;
    cout << s.exist(board, word);
    return 0;
}

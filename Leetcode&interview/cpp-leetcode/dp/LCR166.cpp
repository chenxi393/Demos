#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int jewelleryValue(vector<vector<int>>& frame)
    {
        for (int i = 0; i < frame.size(); i++) {
            for (int j = 0; j < frame[0].size(); j++) {
                int cur = 0;
                if (i > 0) {
                    cur = frame[i - 1][j];
                }
                if (j > 0) {
                    cur = max(cur, frame[i][j - 1]);
                }
                frame[i][j] += cur;
            }
        }
        return frame[frame.size() - 1][frame[0].size() - 1];
    }
};

int main()
{

    return 0;
}

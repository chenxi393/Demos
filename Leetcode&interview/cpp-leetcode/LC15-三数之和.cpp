#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    // 判断是否存在三元组和为0
    // 返回的结果不可用重复
    vector<vector<int>> threeSum(vector<int>& nums)
    {
        vector<vector<int>> res;
        sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size() && nums[i] <= 0; i++) {
            int L = i + 1;
            int R = nums.size() - 1;
            while (L < R) {
                if (nums[i] + nums[L] + nums[R] == 0) {
                    res.push_back({ nums[i], nums[L], nums[R] });
                    while (L + 1 < nums.size() && nums[L] == nums[L + 1]) {
                        L++;
                    }
                    while (R + 1 < nums.size() && nums[R] == nums[R - 1]) {
                        R--;
                    }
                    L++;
                    R--;
                } else if (nums[i] + nums[L] + nums[R] < 0) {
                    L++;
                } else {
                    R--;
                }
            }
            while (i + 1 < nums.size() && nums[i] == nums[i + 1]) {
                i++;
            }
        }
        return res;
    }
};
int main()
{
    Solution s;
    vector<int> in;
    in.push_back(-4);
    in.push_back(-2);
    in.push_back(-2);
    in.push_back(-2);
    in.push_back(0);
    in.push_back(1);
    in.push_back(2);
    in.push_back(2);
    in.push_back(2);
    in.push_back(3);
    in.push_back(3);
    in.push_back(4);
    in.push_back(4);
    in.push_back(6);
    in.push_back(6);

    // [-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]

    vector<vector<int>> res = s.threeSum(in);
    for (int i = 0; i < res.size(); i++) {
        for (int j = 0; j < 3; j++) {
            cout << res[i][j] << " ";
        }
        cout << endl;
    }
    return 0;
}

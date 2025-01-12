#include <bits/stdc++.h>
#include <iostream>
#include <stack>
#include <vector>
using namespace std;

class Solution {
public:
    vector<vector<int>> ans;
    vector<vector<int>> permute(vector<int>& nums)
    {

        dfs(nums, 0);
        return ans;
    }

    void dfs(vector<int> nums, int k)
    {
        if (k == nums.size()) {
            ans.push_back(nums);
            return;
        }
        // 1 2 3
        // 1 2 3
        // 2 1 3
        // 2 3 1
        for (int i = k; i < nums.size(); i++) {
            swap(nums.at(i), nums.at(k));
            dfs(nums, k + 1);
            swap(nums.at(i), nums.at(k));
        }
    }
};

#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target)
    {
        unordered_map<int, int> mp;

        for (int i = 0; i < nums.size(); i++) {
            if (mp.find(nums[i]) != mp.end()) {
                return { i, mp[nums[i]] };
            }
            mp[target - nums[i]] = i;
        }
        return {};
    }
};

int main()
{

    return 0;
}

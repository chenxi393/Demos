#include <iostream>
#include <vector>

using namespace std;

//  O(n) 复杂度 归并 然后取中位数
double findmid(vector<int>& nums1, vector<int>& nums2)
{
    int all = nums1.size() + nums2.size();
    vector<int> nums;
    nums.reserve(all);
    int i = 0, j = 0;
    while (i != nums1.size() && j != nums2.size()) {
        if (nums1[i] < nums2[j]) {
            nums.push_back(nums1[i]);
            i++;
        } else {
            nums.push_back(nums2[j]);
            j++;
        }
    }
    if (i == nums1.size()) {
        while (j != nums2.size()) {
            nums.push_back(nums2[j]);
            j++;
        }
    }
    if (j == nums2.size()) {
        while (i != nums1.size()) {
            nums.push_back(nums1[i]);
            i++;
        }
    }
    if (all % 2 == 1) {
        return nums[all / 2];
    } else {
        return (nums[all / 2] + nums[all / 2 - 1]) / 2.0;
    }
}

int main()
{
    vector<int> a;
    vector<int> b;
    a.push_back(1);
    a.push_back(2);
    b.push_back(3);
    b.push_back(4);
    printf("%lf\n", findmid(a, b));
    return 0;
}

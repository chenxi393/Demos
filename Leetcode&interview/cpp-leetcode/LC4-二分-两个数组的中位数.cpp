#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2)
    {
        // 注意杜绝空数组的情况
        // 一个为空直接返回
        if (nums1.size() == 0) {
            if (nums2.size() % 2 == 0) {
                return (nums2[nums2.size() / 2] + nums2[nums2.size() / 2 - 1]) / 2.0;
            }
            return nums2[nums2.size() / 2];
        }
        if (nums2.size() == 0) {
            if (nums1.size() % 2 == 0) {
                return (nums1[nums1.size() / 2] + nums1[nums1.size() / 2 - 1]) / 2.0;
            }
            return nums1[nums1.size() / 2];
        }
        int allSize = nums1.size() + nums2.size();
        if (allSize % 2 == 0) {
            // 需要找allSize/2 和 allSize/2-1
            return (findK(nums1, nums2, allSize / 2 + 1) + findK(nums1, nums2, allSize / 2)) / 2.0;
        }
        // 找allSize/2-1即可
        return findK(nums1, nums2, allSize / 2 + 1);
    }

    int findK(vector<int>& nums1, vector<int>& nums2, int k)
    {
        int i = 0, j = 0;
        int vi, vj;
        while (k > 1) {
            int temp = k / 2 - 1;
            vi = vj = temp;
            // v设置为间隔
            // 防止越界
            if (temp + i >= nums1.size()) {
                vi = nums1.size() - 1 - i;
            }
            if (temp + j >= nums2.size()) {
                vj = nums2.size() - 1 - j;
            }
            // vi vj为要比较的
            if (nums1[vi + i] >= nums2[vj + j]) {
                // 丢弃vj 中位数不在里面
                j = vj + 1 + j;
                k = k - vj - 1;
            } else {
                // 反之
                i = vi + 1 + i;
                k = k - vi - 1;
            }
            // 检查是否有一个数组为空
            if (i >= nums1.size()) {
                return nums2[j + k - 1];
            }
            if (j >= nums2.size()) {
                return nums1[i + k - 1];
            }
        }
        return min(nums1[i], nums2[j]);
    }
};

int main()
{
    vector<int> a;
    vector<int> b;
    a.push_back(2);
    Solution s;
    cout << s.findMedianSortedArrays(a, b) << endl;
    return 0;
}

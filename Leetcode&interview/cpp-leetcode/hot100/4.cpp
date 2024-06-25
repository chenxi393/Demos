#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // 将问题转化为 有序数组 找第K 小的数
    // 例如14个数 找第7小的数 和第8小的数 求平均
    // 13个数 就可以找第7小的数

    // 这和快速选择不一样 快速选择是在无序数组中 找第K大的值
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2)
    {
        int length = nums1.size() + nums2.size();
        // 如果奇数
        if (length % 2) {
            return findK(nums1, nums2, 0, nums1.size() - 1, 0, nums2.size() - 1, length / 2 + 1);
        }
        return (findK(nums1, nums2, 0, nums1.size() - 1, 0, nums2.size() - 1, length / 2)
                   + findK(nums1, nums2, 0, nums1.size() - 1, 0, nums2.size() - 1, length / 2 + 1))
            * 0.5;
    }

    // 每次去除2/k个数
    int findK(vector<int>& nums1, vector<int>& nums2, int l1, int r1, int l2, int r2, int k)
    {
        if (l1 == nums1.size()) {
            return nums2[l2 + k - 1];
        }
        if (l2 == nums2.size()) {
            return nums1[l1 + k - 1];
        }
        if (k == 1) {
            return min(nums1[l1], nums2[l2]);
        }
        int len1 = r1 - l1 + 1;
        int len2 = r2 - l2 + 1;
        int index1 = r1;
        int index2 = r2;
        if (len1 > k / 2) { // 够长 取2/k的数
            index1 = l1 + k / 2 - 1;
        }
        if (len2 > k / 2) {
            index2 = l2 + k / 2 - 1;
        }

        if (nums1[index1] >= nums2[index2]) {
            return findK(nums1, nums2, l1, r1, index2 + 1, r2, k - (index2 - l2 + 1));
        }
        return findK(nums1, nums2, index1 + 1, r1, l2, r2, k - (index1 - l1 + 1));
    }
};

int main()
{

    return 0;
}
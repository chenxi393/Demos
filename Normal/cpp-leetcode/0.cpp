#include <iostream>
#include <vector>
using namespace std;

class MedianFinder {
private:
    vector<int> nums;

public:
    /** initialize your data structure here. */
    MedianFinder()
    {
    }
    void addNum(int num)
    {
        nums.push_back(num);
    }

    double findMedian()
    {
        if (nums.size() == 0) {
            return 0;
        }
        if (nums.size() % 2 == 1) {
            return quickselect(nums.size() / 2);
        } else {
            return (quickselect(nums.size() / 2) + quickselect(nums.size() / 2 - 1)) / 2.0;
        }
    }
    int quickselect(int k)
    {
        int l = 0, r = nums.size() - 1;
        while (l < r) {
            int i = l, j = r;
            int mid = medianOfThree(l, r, (l + r) / 2);
            swap(nums[r], nums[l]);
            swap(nums[mid], nums[l]);
            int pivot = nums[l];
            while (i < j) {
                while (i < j && nums[j] >= pivot) {
                    j--;
                }
                nums[i] = nums[j];
                while (i < j && nums[i] <= pivot) {
                    i++;
                }
                nums[j] = nums[i];
            }
            nums[i] = pivot;
            if (i == k) {
                break;
            } else if (i < k) {
                l = i + 1;
            } else {
                r = i - 1;
            }
        }
        return nums[k];
    }
    int medianOfThree(int a, int b, int c)
    {
        if ((nums[a] <= nums[b] && nums[b] <= nums[c]) || (nums[c] <= nums[b] && nums[b] <= nums[a])) {
            return b;
        } else if ((nums[b] <= nums[a] && nums[a] <= nums[c]) || (nums[c] <= nums[a] && nums[a] <= nums[b])) {
            return a;
        } else {
            return c;
        }
    }
};

int main()
{

    return 0;
}

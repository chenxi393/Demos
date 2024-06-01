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
            return (quickselect(nums.size() / 2 - 1) + quickselect(0, nums.size() / 2)) / 2.0;
        }
    }
    int quickselect(int k, int l = 0)
    {
        int r = nums.size();
        while (l < r) {
            int i = l, j = r;
            int random = rand() % (r - l + 1) + l;
            int pivot = nums[random];
            nums[random] = nums[i];
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
};

int main()
{

    return 0;
}

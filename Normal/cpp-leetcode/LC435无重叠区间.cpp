#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;
bool cmp(const vector<int>& s1, const vector<int>& s2)
{
    return s1[1] < s2[1];
}

class Solution {
public:
    // 考虑贪心 按区间间隔排序 最小的间隔优先插入 插入不了的丢弃
    // 上面这种想法实际上是错误的
    // 应该贪心的确定最左区间，最左区间的右边界最小
    // 然后继续确定最左区间
    // 当然反过来确定右区间也是可以的
    int eraseOverlapIntervals(vector<vector<int>>& intervals)
    {
        sort(intervals.begin(), intervals.end(), cmp);
        // 最左区间
        int right = intervals[0][1];
        int ans = 0;
        int n = intervals.size();
        for (int i = 1; i < n; i++) {
            if (intervals[i][0] < right) {
                ans++;
            } else {
                right = intervals[i][1];
            }
        }
        return ans;
    }
};

int main()
{

    return 0;
}

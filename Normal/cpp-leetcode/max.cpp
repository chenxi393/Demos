#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

// 只在严格递增严格递减有效
// 前半部分升序后半部分降序，找最大值
// 如果非严格递减 有等于的 那就砍左边一个 或者最右边一个
int findMax(const vector<int>& s, int l, int r)
{
    // 边界处理
    if (l == r || l + 1 == r) {
        return max(s[l], s[r]);
    }
    int m = (l + r) / 2;
    if (s[m] == s[m - 1] && s[m] == s[m + 1]) {
        if (s[l] <= s[m]) {
            return findMax(s, l + 1, r);
        } else {
            return findMax(s, l + 1, r - 1);
        }
    }else if (s[m] > s[m - 1] && s[m] > s[m + 1]) {
        return s[m];
    } else if (s[m] > s[m - 1]) {
    }
    
}

int main()
{

    return 0;
}

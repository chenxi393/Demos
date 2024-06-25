#include <iostream>

using namespace std;
class Solution {
public:
    string longestPalindrome(string s)
    {
        // 枚举奇数串和偶数串
        string res;
        for (int i = 0; i < s.length(); i++) {
            // 奇数串
            int l = i - 1;
            int r = i + 1;
            while (l >= 0 && r < s.length() && s[l] == s[r]) {
                l--;
                r++;
            }
            if (r - l - 1 > res.length()) {
                res = s.substr(l + 1, r - l - 1);
            }

            // 偶数串
            l = i;
            r = i + 1;
            while (l >= 0 && r < s.length() && s[l] == s[r]) {
                l--;
                r++;
            }
            if (r - l - 1 > res.length()) {
                res = s.substr(l + 1, r - l - 1);
            }
        }
        return res;
    }
};
int main()
{

    return 0;
}

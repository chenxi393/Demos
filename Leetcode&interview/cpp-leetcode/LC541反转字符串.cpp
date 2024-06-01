#include <iostream>

using namespace std;
class Solution {
public:
    string reverseStr(string s, int k)
    {
        int f = 0;
        for (int i = 0; i < s.size(); i++) {
            f++;
            if (f == k) {
                int l = i - f + 1;
                int r = i;
                while (l < r) {
                    swap(s[l++], s[r--]);
                }
                f = -f;
            }
        }
        if (f < k) {
            int l = s.size() - 1 - f + 1;
            int r = s.size() - 1;
            while (l < r) {
                swap(s[l++], s[r--]);
            }
        }
        return s;
    }
};

int main()
{

    return 0;
}

#include <iostream>

using namespace std;

class Solution {
public:
    string reverseWords(string s)
    {
        int k = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == ' ') {
                reverse(&s[k], &s[i - 1]);
                k = i + 1;
            }
        }
        reverse(&s[k], &s[s.size() - 1]);
        return s;
    }
    void reverse(char* s1, char* s2)
    {
        while (s1 < s2) {
            swap(*(s1++), *(s2--));
        }
    }
};

int main()
{
    string s = "Let's take LeetCode contest";
    Solution sss;
    cout << sss.reverseWords(s) << endl;
    return 0;
}

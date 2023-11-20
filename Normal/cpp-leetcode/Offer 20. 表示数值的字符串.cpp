#include <iostream>

using namespace std;

///    "^\s*[-+]?(?:\d+\.\d*|\.\d+|\d+)(?:[eE][-+]?\d+)?\s*$"      正则可以试试
class Solution {
public:
    // 小数： 整数+. +无符号整数
    //          整数+.
    //          . 无符号整数
    // 整数

    // 数值
    // 小数e/E 整数
    // 整数e/E 整数
    bool isNumber(string s)
    {
        int i = 0;
        int n = s.length();
        // 1. 去空格
        int p1 = 0, p2 = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] != ' ') {
                p1 = i;
                break;
            }
        }
        for (int i = n - 1; i >= 0; i--) {
            if (s[i] != ' ') {
                p2 = i;
                break;
            }
        }
        string ans = s.substr(p1, p2 - p1 + 1);

        // 2. 看有没有e/E
        int find_e = ans.find('e');
        int find_E = ans.find('E');
        if (find_e == string::npos && find_E == string::npos) {
            return isXiAOSHU(ans);
        } else if (find_e != string::npos && find_E != string::npos) {
            return false;
        } else if (find_e != string::npos) {
            return isXiAOSHU(ans.substr(0, find_e)) && isZHENGSHU(ans.substr(find_e + 1, ans.length() - find_e));
        } else {
            return isXiAOSHU(ans.substr(0, find_E)) && isZHENGSHU(ans.substr(find_E + 1, ans.length() - find_E));
        }
    }

    bool isZHENGSHU(string s, int flag = 1)
    {
        if (flag) {
            if (s[0] != '+' && s[0] != '-') {
                flag = 0;
            }
        }
        if (flag >= s.length()) {
            return false;
        }
        for (int i = flag; i < s.length(); i++) {
            if (s[i] < '0' || s[i] > '9') {
                return false;
            }
        }
        return true;
    }

    bool isXiAOSHU(string s)
    {
        // 整数我们页认为是小数 待判断
        int n = s.length();
        int t = 0;
        if (n > 0 && (s[0] == '+' || s[0] == '-')) {
            t = 1;
        }
        for (int i = t; i < n; i++) {
            if (s[i] == '.') {
                if (i == t) {
                    return isZHENGSHU(s.substr(i + 1, n - i - 1), 0);
                } else if (i == n - 1) {
                    return isZHENGSHU(s.substr(t, n - t - 1), 0);
                } else {
                    return isZHENGSHU(s.substr(t, i - t), 0) && isZHENGSHU(s.substr(i + 1, n - i - 1), 0);
                }
            }
        }
        return isZHENGSHU(s);
    }
};

int main()
{
    Solution s;
    string ttt[20] = {
        "+100",
        "5e2",
        "-123",
        "3.1416",
        "-1E-16",
        "0123",
        "12e",
        "1a3.14",
        "1.2.3",
        "+-5",
        "12e+5.4",
        "    .1  ",
        ".",
        "e",
        "e9"
    };
    for (int i = 0; i < 15; i++) {
        cout << s.isNumber(ttt[i]) << endl;
    }

    return 0;
}

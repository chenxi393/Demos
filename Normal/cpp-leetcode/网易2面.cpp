#include <algorithm>
#include <queue>
#include <vector>
using namespace std;
/*
主持人调度（二）
有 n 个活动即将举办，每个活动都有开始时间与活动的结束时间，第 i 个活动的开始时间是 starti ,第 i 个活动的结束时间是 endi ,举办某个活动就需要为该活动准备一个活动主持人。

一位活动主持人在同一时间只能参与一个活动。并且活动主持人需要全程参与活动，换句话说，一个主持人参与了第 i 个活动，那么该主持人在 (starti,endi) 这个时间段不能参与其他任何活动。求为了成功举办这 n 个活动，最少需要多少名主持人。

数据范围:  ， -2^{32} \le start_i\le end_i \le 2^{31}-1

复杂度要求：时间复杂度  ，空间复杂度
示例 1

输入
2,[[1,2],[2,3]]
输出
1
说明
只需要一个主持人就能成功举办这两个活动
示例 2

输入
2,[[1,3],[2,4]]
输出
2
说明
需要两个主持人才能成功举办这两个活动
*/
bool cmp(vector<int> s1, vector<int> s2)
{
    if (s1[1] == s2[1]) {
        return s1[0] < s2[0];
    }
    return s1[1] < s2[1];
}

class Solution {
public:
    /**
     * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
     *
     * 计算成功举办活动需要多少名主持人
     * @param n int整型 有n个活动
     * @param startEnd int整型vector<vector<>> startEnd[i][0]用于表示第i个活动的开始时间，startEnd[i][1]表示第i个活动的结束时间
     * @return int整型
     */
    //

    /*
    10,
     [-2147483648,-2147483647]
    [-2147483648,-2147483647]
    [-2147483648,-2147483647]
    [-2147483648,-2147483647]]
    ,[-2147483648,-2147483647]
    [2147483646,2147483647],
    [2147483646,2147483647]
    ,[2147483646,2147483647]
    ,[2147483646,2147483647],
     [[2147483646,2147483647]
    https://github.com/codecrafters-io/build-your-own-x
    */
    //
    int minmumNumberOfHost(int n, vector<vector<int>>& startEnd)
    {
        // write code here

        sort(startEnd.begin(), startEnd.end(), cmp);
        priority_queue<int> people;
        people.push(startEnd[0][1]);
        for (int i = 1; i < startEnd.size(); i++) {
            if (people.top() > startEnd[i][0]) {
            } else {
                people.pop();
            }
            people.push(startEnd[i][0]);
        }
        return people.size();
    }
};

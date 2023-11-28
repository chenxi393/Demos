#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    int val;
    ListNode* next;
    ListNode()
        : val(0)
        , next(nullptr)
    {
    }
    ListNode(int x)
        : val(x)
        , next(nullptr)
    {
    }
    ListNode(int x, ListNode* next)
        : val(x)
        , next(next)
    {
    }
};

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    // 不要跳进递归，而是利用明确的定义来实现算法逻辑！ 大赞！
    // 明确的定义是递归的输入输出
    ListNode* reverseBetween(ListNode* head, int left, int right)
    {
        if (left == 1) {
            return reverseListByN(head, right);
        }
        head->next = reverseBetween(head->next, left - 1, right - 1);
        return head;
    }
    ListNode* reverseListByN(ListNode* head, int n)
    {
        if (n == 1) {
            return head;
        }
        // 递归就是比较语义化的
        // reverseList 把后面的都反转好了
        // 返回反转好的链表
        ListNode* finsh = reverseListByN(head->next, n - 1);
        // 保存不需要逆转的
        ListNode* next_ = head->next->next;
        head->next->next = head;
        head->next = next_;
        return finsh;
    }
};

int main()
{

    return 0;
}
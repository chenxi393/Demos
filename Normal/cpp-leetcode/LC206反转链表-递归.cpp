#include <iostream>

using namespace std;
struct ListNode {
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

class Solution {
public:
    // 递归解法
    // 迭代非递归 头插即可
    ListNode* reverseList(ListNode* head)
    {
        if (!head || head->next == NULL) {
            return head;
        }
        // 递归就是比较语义化的
        // reverseList 把后面的都反转好了
        // 返回反转好的链表
        ListNode* finsh = reverseList(head->next);
        // 要把当前节点加到反转好的链表上
        // 那必然当前节点的下一个是尾巴
        head->next->next = head;
        head->next = NULL;
        return finsh;
    }
};

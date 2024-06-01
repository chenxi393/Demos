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

class Solution1 {
public:
    // 12/11 再写一次 反转链表
    ListNode* reverseBetween(ListNode* head, int left, int right)
    {

        if (left == 1) {
            return reverse(head, right);
        }
        head->next = reverseBetween(head->next, left - 1, right - 1);
        return head;
    }
    //
    ListNode* reverse(ListNode* head, int right)
    {
        // head->next 其实就是反转好的链表的尾部
        // 把自己放入尾部 然后返回刚刚的头部即可
        if (right == 1) {
            return head;
        }
        ListNode* res = reverse(head->next, right - 1);
        ListNode* temp = head->next->next;
        head->next->next = head;
        head->next = temp;
        return res;
    }
};
int main()
{
    ListNode L1;
    ListNode L2;
    ListNode L3;
    ListNode L4;
    ListNode L5;
    L1.next = &L2;
    L1.val = 1;
    L2.next = &L3;
    L2.val = 2;
    L3.next = &L4;
    L3.val = 3;
    L4.next = &L5;
    L4.val = 4;
    L5.next = NULL;
    L5.val = 5;

    Solution1 s;
    s.reverseBetween(&L1, 2, 4);

    return 0;
}
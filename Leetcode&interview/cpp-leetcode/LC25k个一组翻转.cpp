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
    ListNode* reverseKGroup(ListNode* head, int k)
    {
        ListNode hh;
        hh.next = head;
        head = &hh;
        ListNode* p = head;
        while (p->next) {
            // _ -> 1 -> 2 -> 3-> 4 -> 5
            ListNode* nn = p;
            for (int i = 0; i < k; i++) {
                if (nn->next == NULL) {
                    return head->next;
                }
                nn = nn->next;
            }
            // tempNext保存新的尾部的下一个
            ListNode* tempNext = nn->next;
            // 返回新的头部
            ListNode* temp = reverse(p->next, tempNext);
            //  nn已不是尾部了 变化了
            //  p->next 会被反转到尾部 实际的尾部应该是 p->next
            p->next->next = tempNext; // 真实的尾部拼接
            // 保存下一轮的p
            ListNode* tempP = p->next;
            p->next = temp; // 改变当前的下一个
            p = tempP;
        }
        return head->next;
    }
    ListNode* reverse(ListNode* head, ListNode* tail_next)
    {
        ListNode hh;
        ListNode* vhh = &hh;
        while (head != tail_next) {
            ListNode* temp = head->next;
            head->next = vhh->next;
            vhh->next = head;
            head = temp;
        }
        return vhh->next;
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

    Solution s;
    s.reverseKGroup(&L1, 2);
    return 0;
}

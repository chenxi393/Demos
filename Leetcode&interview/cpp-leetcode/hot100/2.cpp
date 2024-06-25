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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2)
    {
        ListNode* head = new (ListNode)(-1);
        ListNode* temp = head;
        int t = 0;
        while (l1 != nullptr || l2 != nullptr || t != 0) {
            t += l1 ? l1->val : 0;
            t += l2 ? l2->val : 0;
            ListNode* node = new (ListNode)(t % 10);
            temp->next = node;
            temp = temp->next;
            t /= 10;

            l1 = l1 ? l1->next : nullptr;
            l2 = l2 ? l2->next : nullptr;
        }
        return head->next;
    }
};
int main()
{

    return 0;
}

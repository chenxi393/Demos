#include <iostream>
#include <vector>
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
    ListNode* mergeKLists(vector<ListNode*>& lists)
    {
        if (lists.size() == 0) {
            return NULL;
        }
        return merge(lists, 0, lists.size() - 1);
    }
    ListNode* merge(vector<ListNode*>& lists, int l, int r)
    {
        if (l == r) {
            return lists[l];
        }
        int m = (l + r) / 2;
        ListNode* left = merge(lists, l, m);
        ListNode* right = merge(lists, m + 1, r);

        if (left == NULL) {
            return right;
        }
        if (right == NULL) {
            return left;
        }
        ListNode* ansHead;
        if (left->val < right->val) {
            ansHead = left;
            left = left->next;
        } else {
            ansHead = right;
            right = right->next;
        }
        ListNode* cur = ansHead;
        while (left && right) {
            if (left->val < right->val) {
                cur->next = left;
                left = left->next;
            } else {
                cur->next = right;
                right = right->next;
            }
            cur = cur->next;
        }
        if (left) {
            cur->next = left;
        }
        if (right) {
            cur->next = right;
        }
        return ansHead;
    }
};

int main()
{
    ListNode a1(1);
    ListNode a2(4);
    ListNode a3(5);
    a1.next = &a2;
    a2.next = &a3;
    ListNode b1(1);
    ListNode b2(3);
    ListNode b3(4);
    b1.next = &b2;
    b2.next = &b3;
    ListNode c1(2);
    ListNode c2(6);
    c1.next = &c2;
    Solution s;
    vector<ListNode*> all;
    all.push_back(&a1);
    all.push_back(&b1);
    all.push_back(&c1);
    s.mergeKLists(all);
    return 0;
}

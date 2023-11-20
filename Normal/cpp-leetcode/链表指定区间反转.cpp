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
    ListNode* reverseBetween(ListNode* head, int left, int right)
    {
        ListNode* p1 = head;
        
    }
    ListNode* reverseList(ListNode* head)
    {
        ListNode ans;
        while (head) {
            ListNode* temp = head;
            head = head->next;
            temp->next = ans.next;
            ans.next = temp;
        }
        return ans.next;
    }
};

int main()
{

    return 0;
}
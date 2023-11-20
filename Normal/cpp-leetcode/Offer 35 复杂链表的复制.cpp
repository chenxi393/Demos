#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

// /**
//  * Definition for singly-linked list.
//  * struct ListNode {
//  *     int val;
//  *     ListNode *next;
//  *     ListNode(int x) : val(x), next(NULL) {}
//  * };
//  */

class Node {
public:
    int val;
    Node* next;
    Node* random;

    Node(int _val)
    {
        val = _val;
        next = NULL;
        random = NULL;
    }
};

class Solution {
public:
    Node* copyRandomList(Node* head)
    {
        //    7 ->
        unordered_map<Node*, Node*> mm;
        Node* p = NULL;
        Node* head_2 = NULL;
        while (head) {
            Node* t = new Node(head->val);
            t->random = head->random;
            t->next = NULL;
            if (p == NULL) {
                p = t;
                head_2 = t;

            } else {
                p->next = t;
                p = t;
            }

            mm[head] = t;
            head = head->next;
        }
        p = head_2;
        while (p) {
            p->random = mm[p->random];
            p = p->next;
        }
        return head_2;
    }
};

int main()
{
    Solution s;
    Node* head = new Node(3);
    Node* t1 = new Node(3);
    Node* t2 = new Node(3);
    unordered_map<Node*, Node*> mm;
    if(mm[NULL]==NULL){
        cout<<324234<<endl;
    }
    head->next = t1;
    head->random = NULL;
    t1->next = t2;
    t1->random = head;
    t2->next = NULL;
    t2->random = NULL;
    cout << s.copyRandomList(head);

    return 0;
}

#include <iostream>
#include <unordered_set>
using namespace std;
struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x)
        : val(x)
        , next(NULL)
    {
    }
};
class Solution {
public:
    ListNode* detectCycle(ListNode* head)
    {
        // 直接哈希表 再次碰到即为入环处
        unordered_set<ListNode*> ss;
        while (head) {
            if (ss.find(head) != ss.end()) {
                // 说明找到了
                return head;
            }
            ss.insert(head);
            head = head->next;
        }
        return nullptr;
    }
};

class Solution1 {
public:
    // 快慢指针 相遇点
    //  当相遇时 再来一个指针指向head
    //  当slow和head同时走 再次相遇即为入口点
    // 这是公式计算出的
    ListNode* detectCycle(ListNode* head)
    {
        ListNode* fast = head;
        ListNode* slow = head;
        do {
            if (fast == nullptr || fast->next == nullptr) {
                return nullptr;
            }
            fast = fast->next->next;
            slow = slow->next;
        } while (slow != fast);
        while (slow != head) {
            slow = slow->next;
            head = head->next;
        }
        return head;
    }
};

int main()
{

    return 0;
}

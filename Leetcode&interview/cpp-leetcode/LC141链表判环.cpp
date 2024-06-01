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
    bool hasCycle(ListNode* head)
    {
        unordered_set<ListNode*> ss;
        while (head) {
            if (ss.find(head) != ss.end()) {
                // 说明找到了
                return true;
            }
            ss.insert(head);
            head = head->next;
        }
        return false;
    }
};

// 快慢指针写法
// TODO快慢指针为什么一定会相遇
// 两种情况讨论一下 如果快慢指针相差1
// 各往前走 即相遇
// 如果相差2 各往前走 则相差1 转化为第一种
// 其他距离都可以转化为这两种情况
// 如果快指针走3步可以吗
// 考虑相差1 需要绕一环
// 考虑相差2 刚好一次碰到
// 相差3 就相差2
// 分情况考虑 环为偶数
class Solution1 {
public:
    bool hasCycle(ListNode* head)
    {
        ListNode* fast = head;
        ListNode* slow = head;
        do {
            // 不用判断 slow有没有越界 因为fast一定更快
            if (!fast || !fast->next) {
                return false;
            }
            slow = slow->next;
            fast = fast->next->next;
        } while (slow != fast);
        return true;
    }
};

int main()
{

    return 0;
}

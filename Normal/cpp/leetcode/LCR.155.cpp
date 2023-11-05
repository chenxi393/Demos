#include <iostream>
#include <vector>
using namespace std;
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() { }

    Node(int _val)
    {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right)
    {
        val = _val;
        left = _left;
        right = _right;
    }
};
// 看一下动图就懂了
// 递归是可以的 原采用保存到vetcor里 ---这样肯定会被面试官说。。。
class Solution {
public:
    Node* treeToDoublyList(Node* root)
    {
        if (root == NULL) {
            return NULL;
        }
        dfs(root);
        head->left = pre;
        pre->right = head;
        return head;
    }
    Node *head = NULL, *pre = NULL;
    void dfs(Node* cur)
    {
        if (cur == NULL) {
            return;
        }
        dfs(cur->left);
        if (head == NULL) { // 保存头结点 第一个访问的必是头结点
            head = cur;
        } else {
            pre->right = cur;
            cur->left = pre;
        }
        pre = cur;
        dfs(cur->right);
    }
};
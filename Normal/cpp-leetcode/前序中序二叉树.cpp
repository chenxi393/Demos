#include <iostream>
#include <vector>
using namespace std;
struct TreeNode {
    int data;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int d)
    {
        data = d;
        left = right = NULL;
    }
};
// 前序 和中序遍历 构建二叉树
// 前序 1,2,4,7,3,5,6,8  根左右
// 中序 4,7,2,1,5,3,8,6  左根右

TreeNode* BuildTreeWithFrontAndMid(const vector<int>& front, const vector<int>& mid)
{
    // 为空
    if (front.size() == 0 || mid.size() == 0) {
        return NULL;
    }
    // 前序 的为根节点
    TreeNode* Root = new TreeNode(front[0]);
    // 在中序遍历中找到根节点
    int i = 0;
    for (; i < mid.size(); i++) {
        if (mid[i] == front[0]) {
            break;
        }
    }
    // 要是找不到记得报错
    Root->left = BuildTreeWithFrontAndMid(vector<int>(front.begin() + 1, front.begin() + i + 1), vector<int>(mid.begin(), mid.begin() + i));
    Root->right = BuildTreeWithFrontAndMid(vector<int>(front.begin() + i + 1, front.end()), vector<int>(mid.begin() + i + 1, mid.end()));
    return Root;
}

// 后序 7 4 2 5 8 6 3 1  左右根
// 中序 4,7,2,1,5,3,8,6  左根右
TreeNode* BuildTreeWithMidAndLast(const vector<int>& last, const vector<int>& mid)
{
    // 为空
    if (last.size() == 0 || mid.size() == 0) {
        return NULL;
    }
    // 后序尾 的为根节点
    TreeNode* Root = new TreeNode(last[last.size() - 1]);

    // 在中序遍历中找到根节点
    int i = 0;
    for (; i < mid.size(); i++) {
        if (mid[i] == last[last.size() - 1]) {
            break;
        }
    }
    // 要是找不到记得报错
    Root->left = BuildTreeWithMidAndLast(vector<int>(last.begin(), last.begin() + i), vector<int>(mid.begin(), mid.begin() + i));
    Root->right = BuildTreeWithMidAndLast(vector<int>(last.begin() + i, last.end() - 1), vector<int>(mid.begin() + i + 1, mid.end()));
    return Root;
}

// 后序遍历  左右根
void lastTran(TreeNode* root)
{
    if (root == NULL) {
        return;
    }
    lastTran(root->left);
    lastTran(root->right);
    printf("%d ", root->data);
}
int main()
{
    vector<int> front, mid, end;
    front.push_back(1);
    front.push_back(2);
    front.push_back(4);
    front.push_back(7);
    front.push_back(3);
    front.push_back(5);
    front.push_back(6);
    front.push_back(8);

    mid.push_back(4);
    mid.push_back(7);
    mid.push_back(2);
    mid.push_back(1);
    mid.push_back(5);
    mid.push_back(3);
    mid.push_back(8);
    mid.push_back(6);

    end.push_back(7);
    end.push_back(4);
    end.push_back(2);
    end.push_back(5);
    end.push_back(8);
    end.push_back(6);
    end.push_back(3);
    end.push_back(1);

    auto ans1 = BuildTreeWithFrontAndMid(front, mid);
    lastTran(ans1);
    cout << endl;

    auto ans2 = BuildTreeWithMidAndLast(end, mid);
    lastTran(ans2);
    cout << endl;
    return 0;
}

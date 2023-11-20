#include <iostream>
using namespace std;
// 改成位运算 感觉也没差太多
const int MAX_SIZE = 769;
class MyHashMap {
private:
    struct Hnode {
        Hnode* next;
        int data;
        int key;
    };
    Hnode* map[MAX_SIZE];

public:
    MyHashMap()
    {
        // 改造成头结点 不用在意特殊情况
        for (int i = 0; i < MAX_SIZE; i++) {
            map[i] = new Hnode;
            map[i]->key = -1;
            map[i]->next = NULL;
        }
    }

    void put(int key, int value)
    {
        int hash = key % MAX_SIZE;
        Hnode* head = map[hash]->next;
        while (head) {
            if (head->key == key) {
                head->data = value;
                return;
            }
            head = head->next;
        }
        Hnode* t = new Hnode;
        t->data = value;
        t->key = key;
        t->next = map[hash]->next;
        map[hash]->next = t;
    }

    void remove(int key)
    {
        int hash = key % MAX_SIZE;
        Hnode* head = map[hash];
        Hnode* next = head->next;
        while (next) {
            if (next->key == key) {
                head->next = next->next;
                delete next;
                return;
            }
            head = next;
            next = next->next;
        }
    }

    int get(int key)
    {
        int hash = key % MAX_SIZE;
        Hnode* head = map[hash]->next;
        while (head) {
            if (head->key == key) {
                return head->data;
            }
            head = head->next;
        }
        return -1;
    }
};

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet* obj = new MyHashSet();
 * obj->add(key);
 * obj->remove(key);
 * bool param_3 = obj->contains(key);
 */

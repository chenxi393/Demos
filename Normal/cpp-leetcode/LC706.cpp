// 没有头结点 但非常高效的写法
// 重点是remove
const int MAX_SIZE = 769;
class MyHashMap {
private:
    struct Hnode {
        Hnode* next;
        int data;
        int key;
    };
    Hnode* map[MAX_SIZE] = { nullptr };

public:
    MyHashMap()
    {
    }

    void put(int key, int value)
    {
        int hash = key % MAX_SIZE;
        Hnode* head = map[hash];
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
        t->next = map[hash];
        map[hash] = t;
    }

    void remove(int key)
    {
        int hash = key % MAX_SIZE;
        Hnode** head = &map[hash];
        while (*head) {
            if ((*head)->key == key) {
                Hnode* temp = (*head);// *head 就是真正保存了new出来的那一块地址
                *head = (*head)->next;// *head的值已经保存到temp 把head保存的地址变为下一个 
                delete temp; // 删除new出来的
                return;
            }
            head = &(*head)->next; // 这里取的是（指针保存的结点实体）的next域的地址
        }
    }

    int get(int key)
    {
        int hash = key % MAX_SIZE;
        Hnode* head = map[hash];
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
 * Your MyHashMap object will be instantiated and called as such:
 * MyHashMap* obj = new MyHashMap();
 * obj->put(key,value);
 * int param_2 = obj->get(key);
 * obj->remove(key);
 */
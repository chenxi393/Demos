// #pragma pack (4)
#include <iostream>
#include <vector>

using namespace std;
#define Myoffsetof(type, member) ((size_t)(&((type*)0)->member))

// 内存对齐问题
struct test {
    int a;
    char b[5]; // 好像会有内存对齐的问题 实际上5个字节会占用8个字节
    int c;
};

struct S2
{
    int i:8;    
    char j:4;  
    double b; // 结构体总体大小能够被最宽的成员的大小整除，如不能则在后面补充字节
    int a:4;  
};

struct S3
{
    int i;    
    char j;  
    double b;
    int a;     
};

class A {
public:
    A()
    {
        mX = new int();
    }
    ~A()
    {
        if (mX)
            delete mX;
        mX = NULL;
    }
    int* mX;
};
vector<A> gA;
void add()
{
    A a;
    gA.push_back(a);
}

// char* 和string的问题 以及传参的问题
void print(const string& s)
{
    cout << s << endl;
}

// 这样就不可以调用
void print(string& s)
{
    cout << s << endl;
}

int main()
{
    const string s = "hdua";
    const char* abc = "abc";
    //*abc = 'd';
    print(abc);

    cout << Myoffsetof(test, c) << endl;
    cout << offsetof(test, c) << endl;

    printf("%lu\n", sizeof(S3)); 
    return 0;
}

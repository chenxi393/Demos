#include <iostream>
#include <vector>
using namespace std;
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

int main()
{
    add();
    return 0;
}

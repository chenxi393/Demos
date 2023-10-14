#include <iostream>
#include <map>
using namespace std;

int main()
{
    map<int, int*> mm;
    for (int i = 0; i < 4; i++) {
        mm[i] = &i;
    }

    for (auto t : mm) {
        printf(" key = %d address = %p value = %d\n", t.first, t.second, *t.second);
    }

    return 0;
}

#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int candy(vector<int>& ratings)
    {
        vector<int> all(ratings.size(), 1);
        for (int i = 1; i < ratings.size(); i++) {
            if (ratings[i - 1] > ratings[i] && all[i - 1] <= all[i]) {
                all[i - 1] = all[i] + 1;
            } else if (ratings[i - 1] < ratings[i] && all[i - 1] >= all[i]) {
                all[i] = all[i - 1] + 1;
            }
        }
        for (int i = ratings.size() - 1; i > 0; i--) {
            if (ratings[i - 1] > ratings[i] && all[i - 1] <= all[i]) {
                all[i - 1] = all[i] + 1;
            } else if (ratings[i - 1] < ratings[i] && all[i - 1] >= all[i]) {
                all[i] = all[i - 1] + 1;
            }
        }
        int sum = 0;
        for (int i = 0; i < all.size(); i++) {
            sum += all[i];
        }
        return sum;
    }
};

int main()
{

    return 0;
}

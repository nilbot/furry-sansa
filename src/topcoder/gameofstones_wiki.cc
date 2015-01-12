int count(vector<int> stones)
{
    int s = accumulate(stones.begin(), stones.end(), 0);
    int n = stones.size();
    if (s % n != 0) {
        return -1;
    }
    int r = s / n;
    int cost = 0;
    for (int x: stones) {
        if ( x % 2 != r % 2 ) {
            return -1;
        }
        if (x > r) {
            cost += (x - r) / 2;
        }
    }
    return cost;
}
